package main

import (
	"context"
	"fmt"
	"github.com/zhaoshoucheng/hodgepodge/IoC/app"
	_ "github.com/zhaoshoucheng/hodgepodge/IoC/module"
	_ "github.com/zhaoshoucheng/hodgepodge/IoC/resource"
	_ "github.com/zhaoshoucheng/hodgepodge/IoC/service"
)

func main() {
	err := app.DefaultApplicationLifeCycle().Launch()
	if err != nil {
		fmt.Println("DefaultApplicationLifeCycle Launch err ", err)
		return
	}
	var s1 app.Service1
	app.GetOrCreateRootContainer().Invoke(func(service app.Service1) {
		s1 = service
	})
	s1.AddData("IOC Test")

	app.DefaultApplicationLifeCycle().Shutdown(context.Background())

	//wait
	for i := 0; i < 5; i++ {
		var s2 app.Service2
		app.GetOrCreateRootContainer().Invoke(func(service app.Service2) {
			s2 = service
			initDone, _ := app.DefaultApplicationLifeCycle().RegisterInitializeAwait(fmt.Sprintf("service: %d", i))
			s2.InitService(initDone)
		})
		s2.SyncData(i)
	}
	app.DefaultApplicationLifeCycle().WaitUntilInitialized(context.Background())
	fmt.Println("App Wait Over")
}
