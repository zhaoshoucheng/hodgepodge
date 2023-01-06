package resource

import (
	"context"
	"fmt"
	"github.com/berkaroad/ioc"
	"github.com/zhaoshoucheng/hodgepodge/IoC/app"
)

type ResourceObj struct {
	name string
}

func (r *ResourceObj) Save(str string) {
	fmt.Println(r.name, " Save ", str)
}
func (r *ResourceObj) Remove(str string) {
	fmt.Println(r.name, " Remove ", str)
}

func init() {
	app.DefaultApplicationLifeCycle().RegisterInitializer("resource", func() error {
		mo := &ResourceObj{name: "mongo"}
		// static assert 静态断言类型检测
		func(t app.Resource) {}(mo)
		//rd := &ResourceObj{name: "redis"}
		app.GetOrCreateRootContainer().RegisterTo(mo, (*app.Resource)(nil), ioc.Singleton)
		//app.GetOrCreateRootContainer().RegisterTo(rd, (*app.Resource)(nil), ioc.Singleton)
		return nil
	}, 1)
	app.DefaultApplicationLifeCycle().RegisterFinalizer("resource", func(ctx context.Context) error {
		fmt.Println("resource exit")
		return nil
	}, 1)
}
