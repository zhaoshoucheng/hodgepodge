package resource

import (
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
	mo := &ResourceObj{name: "mongo"}
	// static assert 静态断言类型检测
	func(t app.Resource) {}(mo)
	//rd := &ResourceObj{name: "redis"}
	app.GetOrCreateRootContainer().RegisterTo(mo, (*app.Resource)(nil), ioc.Singleton)
	//app.GetOrCreateRootContainer().RegisterTo(rd, (*app.Resource)(nil), ioc.Singleton)
}
