package service

import (
	"context"
	"fmt"
	"github.com/berkaroad/ioc"
	"github.com/zhaoshoucheng/hodgepodge/IoC/app"
)

var (
	module app.Module

	service2 app.Service2
)

type Service1 struct {
}

func (s1 *Service1) AddData(str string) {
	service2.AddData(str)
	fmt.Println("Service1 AddData ", str)
	module.DataToSave(str)
}
func (s1 *Service1) DelData(str string) {
	service2.DelData(str)
	fmt.Println("Service1 DelData ", str)
	module.DataToRemove(str)
}

func init() {
	app.DefaultApplicationLifeCycle().RegisterInitializer("service", func() error {

		s1 := &Service1{}
		s2 := &Service2{}

		service2 = s2

		//static assert 静态断言做类型检查
		func(t app.Service1) {}(s1)
		func(t app.Service2) {}(s2)

		app.GetOrCreateRootContainer().RegisterTo(s1, (*app.Service1)(nil), ioc.Singleton)
		app.GetOrCreateRootContainer().RegisterTo(s2, (*app.Service2)(nil), ioc.Singleton)

		app.GetOrCreateRootContainer().Invoke(func(mod app.Module) {
			module = mod
		})
		return nil
	}, 3)

	app.DefaultApplicationLifeCycle().RegisterFinalizer("service", func(ctx context.Context) error {
		fmt.Println("service exit")
		return nil
	}, 3)
}
