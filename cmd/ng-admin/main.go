package main

import (
	"context"
	"gin-ng-admin/internal/admin"
	"gin-ng-admin/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
)

var VERSION = "0.0.0"

func main() {
	logger.SetVersion(VERSION)
	ctx := logger.NewTagContext(context.Background(), "__main__")
	app := cli.NewApp()
	app.Name = "gin-ng-admin"
	app.Version = VERSION
	app.Usage = "GIN + NG12 + GORM + CASBIN + WIRE."
	app.Commands = []*cli.Command{
		newCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
		return
	}
}

func newCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "运行web服务",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "配置文件(.yaml)",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "model",
				Aliases:  []string{"m"},
				Usage:    "casbin的访问控制模型(.conf)",
				Required: false,
			},
			&cli.StringFlag{
				Name:  "menu",
				Usage: "初始化菜单数据配置文件(.yaml)",
			},
			&cli.StringFlag{
				Name:  "www",
				Usage: "静态站点目录",
			},
		},
		Action: func(c *cli.Context) error {
			return admin.Run(ctx,
				admin.SetConfigFile(c.String("config")))
		},
	}
}
