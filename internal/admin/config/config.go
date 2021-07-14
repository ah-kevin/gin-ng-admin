package config

import (
	"flag"
	"fmt"
	"gin-ng-admin/pkg/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var (
	C Server
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", utils.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GLOBAL_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config) // 指定配置文件路径
	v.SetConfigType("yaml")
	err := v.ReadInConfig() //读取配置信息
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 监控配置文件变化
	v.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量C
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&C); err != nil {
			fmt.Println(err)
			return
		}
	})
	// 将读取的配置信息保存至全局变量C
	if err := v.Unmarshal(&C); err != nil {
		fmt.Println(err)
	}
	return v
}

type Server struct {
	RunMode     string `json:"run-mode,omitempty;" mapstructure:"run-mode"`
	WWW         string `json:"www,omitempty" mapstructure:"www"`
	Swagger     bool   `json:"swagger,omitempty" mapstructure:"swagger"`
	PrintConfig bool   `json:"print-config,omitempty" mapstructure:"print-config"`
	Log         Log    `json:"log" mapstructure:"log"`
}
