package main

import (
	"flag"
	"net/http"

	"github.com/WechatWall/common"
	"github.com/WechatWall/handler"
)

func main() {
	configPath := flag.String("config", "", "config file's path")
	flag.Parse()

	common.InitConfig(*configPath)

	if err := http.ListenAndServe(common.Config.Listen, handler.Wx()); err != nil {
		panic(err.Error())
	}
	//if err := http.ListenAndServe(":80", handler.Wx()); err != nil {
	//	panic(err.Error())
	//}
}
