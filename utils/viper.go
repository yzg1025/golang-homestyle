package utils

import (
	"fmt"
	"gin/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func ViperTool() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("./config/")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig();err != nil {
		fmt.Printf("err:%s\n",err)
	}

	if err := v.Unmarshal(&global.GCONFIG) ; err != nil{
		fmt.Printf("err:%s",err)
	}

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GCONFIG) ; err != nil{
			fmt.Printf("err:%s",err)
		}
	})

	v.WatchConfig()
}