package main

import (
	"flag"
	"gin-micro/protos/user"
	"gin-micro/user/services"
	"github.com/hero1s/golib/log"
	smicro "github.com/hero1s/golib/micro"
	"github.com/hero1s/golib/utils"
	"github.com/micro/go-micro/v2"
	"os"
	"strings"
)

var (
	etcdAddr = flag.String("etcd", "172.16.3.21:2379", "register etcd address")
	etcdUser = flag.String("user", "", "etcd username")
	etcdPass = flag.String("pass", "", "etcd password")
	confName = flag.String("conf", "", "")
)

func main() {
	flag.Parse()
	smicro.InitGlobalTracer(smicro.JaegerConf{ServiceName: "tracer-srv", Addr: "172.16.3.21:6831"})

	smicro.InitServer("micro.user.svr", "1.0.1", ":8080", &smicro.EtcdRegistry{
		Addrs: strings.Split(*etcdAddr, ","),
		User:  *etcdUser,
		Pass:  *etcdPass,
	}, nil, func(s micro.Service) {
		s.Init(
			micro.BeforeStart(func() error {
				log.Debug("启动前的日志打印")
				return nil
			}),
			micro.AfterStart(func() error {
				log.Debug("启动后的日志打印")
				return nil
			}),
			micro.BeforeStop(func() error {
				log.Debug("服务停止前")
				return nil
			}),
			micro.AfterStop(func() error {
				log.Debug("服务停止后")
				return nil
			}),
		)

		// 注册所有的Handler
		err := user.RegisterUserServiceHandler(s.Server(), new(services.UserService))
		if err != nil {
			log.Error("handler注册失败：" + err.Error())
			os.Exit(0)
		}
	})

	utils.RunMain(func() error {
		return nil
	}, func() {

	}, "pid.txt")

}
