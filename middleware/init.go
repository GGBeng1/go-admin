package middleware

import (
	"fmt"
	"hello/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

func init() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.Server); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.Server); err != nil {
		fmt.Println(err)
	}
	global.Conf = v
}
