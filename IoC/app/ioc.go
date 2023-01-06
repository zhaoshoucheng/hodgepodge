package app

import (
	"context"
	"fmt"
	"github.com/berkaroad/ioc"
	"github.com/spf13/viper"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// 用viper 实现的一个单例
func GetOrCreateRootContainer() ioc.Container {
	v := viper.Get("runtime.container")
	if v == nil {
		v = ioc.NewContainer()
		viper.Set("runtime.container", v)
	}
	return v.(ioc.Container)
}

var appInit = &ApplicationLifeCycle{}

func DefaultApplicationLifeCycle() *ApplicationLifeCycle {
	return appInit
}

type InitProc func() error
type InitializeDoneProc func(err error)
type ShutdownProc func(ctx context.Context) error

type initProcCtx struct {
	name     string
	priority int
	proc     InitProc
}

type shutdownProcCtx struct {
	name     string
	priority int
	proc     ShutdownProc
}

type ApplicationLifeCycle struct {
	allProcs          []initProcCtx     // 需要初始化执行列表
	allShutdown       []shutdownProcCtx // 需要销毁执行列表
	launched          bool              // 已加载标识
	shutdownRequested bool              // 已退出标识
	served            bool              // 异步函数加载标识
	once              sync.Once         //
	waitCounter       int32
}

/*
根据优先级启动各对象的初始化函数
*/
func (alc *ApplicationLifeCycle) Launch() error {
	if alc.launched {
		return fmt.Errorf("already launched")
	}
	min := math.MaxInt32
	max := math.MinInt32
	for _, ctx := range alc.allProcs {
		if ctx.priority > max {
			max = ctx.priority
		}
		if ctx.priority < min {
			min = ctx.priority
		}
	}

	for i := min; i <= max; i++ {
		for _, ctx := range alc.allProcs {
			if ctx.priority == i {
				fmt.Println("calling initializer for ", ctx.name)
				err := ctx.proc()
				if err != nil {
					return fmt.Errorf("luanch application failed: %v", err)
				}
			}
		}
	}

	alc.launched = true
	return nil
}

/*
根据优先级执行各对象的初始化函数
*/
func (alc *ApplicationLifeCycle) Shutdown(ctx context.Context) {
	alc.shutdownRequested = true
	alc.once.Do(func() {
		min := math.MaxInt32
		max := math.MinInt32
		for _, sctx := range alc.allShutdown {
			if sctx.priority > max {
				max = sctx.priority
			}
			if sctx.priority < min {
				min = sctx.priority
			}
		}

		for i := min; i <= max; i++ {
			for _, sctx := range alc.allShutdown {
				if sctx.priority == i {
					fmt.Println("calling finalizer for ", sctx.name)
					err := sctx.proc(ctx)
					if err != nil {
						fmt.Println("shutdown application failed: ", err)
					}
				}
			}
		}

		fmt.Println("application shutdown")
	})
}

/*
注册初始化函数，本质是一个优先级队列，可以通过优先级控制函数执行顺序，而不是根据根据包的引用顺序。
Launch函数执行
*/
func (alc *ApplicationLifeCycle) RegisterInitializer(name string, proc InitProc, priority int) error {
	if alc.launched {
		return fmt.Errorf("too late to register intializer, application already launched")
	}
	alc.allProcs = append(alc.allProcs, initProcCtx{name, priority, proc})
	return nil
}

/*
注册退出清理函数,Shutdown执行
*/
func (alc *ApplicationLifeCycle) RegisterFinalizer(name string, proc ShutdownProc, priority int) error {
	if alc.shutdownRequested {
		return fmt.Errorf("too late to register finalizer")
	}
	alc.allShutdown = append(alc.allShutdown, shutdownProcCtx{
		name:     name,
		priority: priority,
		proc:     proc,
	})
	return nil
}

/*
用于流程控制，所有异步的初始化完成之后，这里才会跳出循环，利用原子计数器实现的一个自旋锁
*/
func (alc *ApplicationLifeCycle) WaitUntilInitialized(ctx context.Context) {
	alc.served = true
	// alc.serveWaitGroup.Wait()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if alc.waitCounter == 0 {
				return
			}
			time.Sleep(time.Millisecond * 50) // sleep 50ms
		}
	}
}

/*
用于流程控制 原子计数器+1，与-1的处理闭包函数
*/
func (alc *ApplicationLifeCycle) RegisterInitializeAwait(name string) (InitializeDoneProc, error) {
	if alc.served {
		return nil, fmt.Errorf("too late to register intialize await")
	}
	once := sync.Once{}
	atomic.AddInt32(&alc.waitCounter, 1)
	fmt.Println("register intialize await for ", name)
	return func(err error) {
		once.Do(func() {
			if err != nil {
				fmt.Println(fmt.Sprintf("initialize await `%s` report a failure: %v", name, err))
			} else {
				fmt.Println(fmt.Sprintf("`%s` initialized", name))
			}
			atomic.AddInt32(&alc.waitCounter, -1)
		})
	}, nil
}
