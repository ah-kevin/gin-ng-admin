package admin

import (
	"context"
	"fmt"
	"gin-ng-admin/internal/admin/config"
	"gin-ng-admin/internal/admin/global"
)

type options struct {
	ConfigFile string
}

// Option 定义配置项
type Option func(*options)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}
func Run(ctx context.Context, opts ...Option) error {
	global.VP = config.Viper() // 初始化Viper
	fmt.Println(config.C)
	return nil
}
