package module

import (
	"context"
	"fmt"
	"github.com/berkaroad/ioc"
	"github.com/zhaoshoucheng/hodgepodge/IoC/app"
)

var (
	rs app.Resource
)

type ModuleObj struct {
}

func (mo *ModuleObj) DataToSave(str string) {
	fmt.Println("ModuleObj DataToSave ", str)
	rs.Save(str)
}
func (mo *ModuleObj) DataToRemove(str string) {
	fmt.Println("ModuleObj DataToRemove ", str)
	rs.Remove(str)
}

func init() {
	app.DefaultApplicationLifeCycle().RegisterInitializer("module", func() error {
		mo := &ModuleObj{}
		// static assert 静态断言类型检测
		func(t app.Module) {}(mo)

		app.GetOrCreateRootContainer().RegisterTo(mo, (*app.Module)(nil), ioc.Singleton)

		app.GetOrCreateRootContainer().Invoke(func(r app.Resource) {
			rs = r
		})
		return nil
	}, 2)

	app.DefaultApplicationLifeCycle().RegisterFinalizer("module", func(ctx context.Context) error {
		fmt.Println("module exit")
		return nil
	}, 2)
}
