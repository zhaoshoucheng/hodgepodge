package service

import (
	"fmt"
)

type Service2 struct {
}

func (s1 *Service2) AddData(str string) {
	fmt.Println("Service2 AddData ", str)
	module.DataToSave(str)
}
func (s1 *Service2) DelData(str string) {
	fmt.Println("Service2 DelData ", str)
	module.DataToRemove(str)
}
