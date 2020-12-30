package conf

import (
	"github.com/hero1s/golib/cache"
	"github.com/hero1s/golib/conf"
	cconf "github.com/hero1s/golib/connsvr/conf"
	"github.com/hero1s/golib/db"
	"github.com/hero1s/golib/log"
)

type SvrConf struct {
	LogLevel   string `json:"log_level"`
	LogPath    string `json:"log_path" `
	WSAddr     string `json:"ws_addr" `
	TCPAddr    string `json:"tcp_addr" `
	MaxConnNum int    `json:"max_conn_num" `
}

type Conf struct {
	Etcd   []string        `json:"etcd"`
	Mysql  []db.DbConf     `json:"mysql"`
	Redis  cache.RedisConf `json:"redis"`
	Server SvrConf         `json:"server"`
}

var Config Conf

func InitConf(confPath string) bool {

	err := conf.AutoParseFile(confPath, &Config)
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("%+v", Config)
	}
	cconf.ConsolePort = 8080

	return true
}
