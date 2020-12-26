package msg

import (
	"github.com/hero1s/golib/connsvr/network/json"
	"github.com/hero1s/golib/connsvr/network/protobuf"
)

var (
	JSONProcessor     = json.NewProcessor()
	ProtobufProcessor = protobuf.NewProcessor()
)

func init() {
	JSONProcessor.Register(&C2S_UserHeartBeat{})
	JSONProcessor.Register(&C2S_UserLogin{})
	JSONProcessor.Register(&S2C_UserLogin{})
	JSONProcessor.Register(&C2S_Message{})
	JSONProcessor.Register(&S2C_Message{})
}

//心跳包
type C2S_UserHeartBeat struct {
	Ptime int64 `json:"ptime" desc:"当前时间"`
}

//用户登录
type C2S_UserLogin struct {
	Token     string `json:"token" desc:"平台token"`
	LoginPlat int    `json:"login_plat" desc:登录平台`
}

//登录返回
type S2C_UserLogin struct {
	Ret int64  `json:"ret" desc:"登录结果"`
	Msg string `json:"msg" desc:"登录返回"`
}

//用户消息
type C2S_Message struct {
	MsgId   int64       `json:"msg_id" desc:"消息ID"`
	Message interface{} `json:"message" desc:"客户端发送消息到服务器"`
}

//服务器发送客户端消息
type S2C_Message struct {
	Uid     uint64 `json:"uid" desc:"用户ID"`
	Message string `json:"message" desc:"消息内容"`
}

