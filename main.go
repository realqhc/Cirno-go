package main

import (
	"log"
	"os"

	"./ciweimao"
	"./config"
	"./fetch"
	"./structure"
	"./util"
	"github.com/urfave/cli"
)

var AppConfig structure.ConfigStruct

func init() {
	AppConfig = config.Load()
	util.InitReq()
}

func main() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			var args = c.Args()
			switch args.Get(0) {
			case "demo":
				ciweimao.Demo()
			case "login":
				ciweimao.Login(AppConfig)
			case "search":
				ciweimao.Search(args.Get(1), 0, AppConfig)
			case "download":
				fetch.DownloadText(args.Get(1), AppConfig)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
