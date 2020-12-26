package main

import (
	"flag"
	"gin-micro/chat_server/conf"
	"gin-micro/chat_server/game"
	"gin-micro/chat_server/gate"
	"github.com/hero1s/golib/connsvr"
	"github.com/hero1s/golib/helpers/file"
)

var confFile string

func init() {
	flag.StringVar(&confFile, "conf", "conf.toml", "default config")
}

func main() {
	flag.Parse()
	file.WritePidFile("pid.txt")

	if !conf.InitConf(confFile) {
		return
	}
	connsvr.Run(
		game.Module,
		gate.Module,
	)
}
