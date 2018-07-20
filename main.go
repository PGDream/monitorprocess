package main

import (
	"github.com/urfave/cli"
	"errors"
	"os"
	"log"
)

func init(){
	log.SetPrefix("【main】")
}

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "c",
			Usage: "请输入配置文件目录",
		},
	}

	app.Action = func(c *cli.Context) error {
		configFile := c.String("c")
		if configFile == "" {
			return errors.New("配置文件不能为空")
		}
		_, err := os.Open(configFile)
		if err != nil && os.IsNotExist(err) {
			return errors.New("配置文件不存在")
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("获取参数异常,信息: ", err)
	}
}
