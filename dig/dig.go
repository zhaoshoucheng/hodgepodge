package dig

import (
	"fmt"
	"go.uber.org/dig"
)

type Config struct {

}

func NewConfig() *Config {
	fmt.Println("NewConfig")
	return &Config{}
}

type Struct1 struct {
}

func NewStruct1(config *Config) *Struct1{
	fmt.Println("NewStruct1")
	return &Struct1{}
}

type Struct2 struct {
}
func NewStruct2(config *Config) *Struct2{
	fmt.Println("NewStruct2")
	return &Struct2{}
}

type Struct3 struct {
	Name string
}
func NewStruct3(s1 *Struct1, s2 *Struct2) *Struct3{
	fmt.Println("NewStruct3")
	return &Struct3{"This is Struct3"}
}

func BuildContainer() (container *dig.Container,err error) {
	container = dig.New()
	err = container.Provide(NewConfig)
	if err != nil {
		return
	}
	err = container.Provide(NewStruct1)
	if err != nil {
		return
	}
	err = container.Provide(NewStruct2)
	if err != nil {
		return
	}
	err = container.Provide(NewStruct3)
	if err != nil {
		return
	}
	return
}

func GetNewStruct3() (s3 *Struct3,err error) {
	c, err :=  BuildContainer()
	if err != nil {
		return
	}
	err = c.Invoke(func(containerS3 *Struct3) {
		s3 = containerS3
	})
	return
}