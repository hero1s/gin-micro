package gate

import (
	"gin-micro/chat_server/game"
	"gin-micro/chat_server/msg"
	"github.com/hero1s/golib/connsvr/gate"
	"github.com/hero1s/golib/log"
	"time"
)

func init() {
	msg.JSONProcessor.SetRouter(&msg.C2S_UserHeartBeat{}, game.ChanRPC)
	msg.JSONProcessor.SetRouter(&msg.C2S_UserLogin{}, game.ChanRPC)
	msg.JSONProcessor.SetRouter(&msg.C2S_Message{}, game.ChanRPC)
	msg.JSONProcessor.SetHandler(&msg.C2S_UserHeartBeat{}, handleHeartBeat)
}
//测试1
func handleHeartBeat(args []interface{}) {
	m := args[0].(*msg.C2S_UserHeartBeat)
	a := args[1].(gate.Agent)
	log.Infof("gate处理 接受websocket心跳包:%v message:%#v", a.RemoteAddr(), m)

	a.WriteMsg(&msg.C2S_UserHeartBeat{Ptime: time.Now().Unix()})
}