package app

import (
	"github.com/berkaroad/ioc"
	"github.com/spf13/viper"
)

func GetOrCreateRootContainer() ioc.Container {
	v := viper.Get("runtime.container")
	if v == nil {
		v = ioc.NewContainer()
		viper.Set("runtime.container", v)
	}
	return v.(ioc.Container)
}
