package clients

import (
	"gin-micro/protos/user"
	smicro "github.com/hero1s/golib/micro"
	"github.com/micro/go-micro/v2/client"
)

func InitClient(conf smicro.EtcdRegistry) {
	smicro.InitClient("", "1.0.1", &conf, nil, func(c client.Client) {
		Client = c
	})
	UserService = user.NewUserService("micro.user.svr", Client)
}

// 定义 client 对象
var Client client.Client

var UserService user.UserService
