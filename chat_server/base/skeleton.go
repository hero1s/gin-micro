package base

import (
	"gin-micro/chat_server/conf"
	"github.com/hero1s/golib/connsvr/chanrpc"
	"github.com/hero1s/golib/connsvr/module"
)

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              conf.GoLen,
		TimerDispatcherLen: conf.TimerDispatcherLen,
		ChanRPCServer:      chanrpc.NewServer(conf.ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}
