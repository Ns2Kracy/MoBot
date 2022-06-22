package config

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var configFile = "./config.yaml"

//用于加载配置文件
func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 默认值
			config = configFile
			fmt.Printf("您正在使用config的默认值,config的路径为%v\n", configFile)
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config) //加载配置文件
	v.SetConfigType("yaml") //设置配置文件类型 支持 JSON/TOML/YAML/HCL/envfile/Java properties 等多种格式的配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		//装配总配置文件
		if err := v.Unmarshal(&GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	//装配总配置文件
	if err := v.Unmarshal(&GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
