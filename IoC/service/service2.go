package service

import (
	"fmt"
	"github.com/zhaoshoucheng/hodgepodge/IoC/app"
	"time"
)

type Service2 struct {
	initDone app.InitializeDoneProc
}

func (s2 *Service2) InitService(done app.InitializeDoneProc) {
	s2.initDone = done
}

func (s2 *Service2) AddData(str string) {
	fmt.Println("Service2 AddData ", str)
	module.DataToSave(str)
}
func (s2 *Service2) DelData(str string) {
	fmt.Println("Service2 DelData ", str)
	module.DataToRemove(str)
}

func (s2 *Service2) SyncData(t int) {
	time.Sleep(time.Second * 2 * time.Duration(t))
	fmt.Println("SyncData over : ", t)
	if s2.initDone != nil {
		s2.initDone(nil)
	}
}
