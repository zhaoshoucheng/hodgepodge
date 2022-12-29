package main

import (
	"github.com/zhaoshoucheng/hodgepodge/IoC/app"
	_ "github.com/zhaoshoucheng/hodgepodge/IoC/module"
	_ "github.com/zhaoshoucheng/hodgepodge/IoC/resource"
	_ "github.com/zhaoshoucheng/hodgepodge/IoC/service"
)

func main() {
	var s1 app.Service1
	app.GetOrCreateRootContainer().Invoke(func(service app.Service1) {
		s1 = service
	})
	s1.AddData("IOC Test")
}
