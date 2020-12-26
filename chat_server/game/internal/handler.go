package internal

import (
	"gin-micro/chat_server/msg"
	"github.com/hero1s/golib/connsvr/gate"
	"github.com/hero1s/golib/helpers/token"
	"github.com/hero1s/golib/log"
	"github.com/hero1s/golib/stringutils"
	"reflect"
	"time"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.C2S_UserHeartBeat{}, handleHeartBeat)
	handleMsg(&msg.C2S_UserLogin{}, handleUserLogin)
	handleMsg(&msg.C2S_Message{}, handleMessage)
}

func handleHeartBeat(args []interface{}) {
	m := args[0].(*msg.C2S_UserHeartBeat)
	a := args[1].(gate.Agent)
	log.Infof("game 接受websocket心跳包:%v message:%#v", a.RemoteAddr(), m)
	info, ok := a.UserData().(UserInfo)
	if ok && info.Uid != 0 {
		info.Ptime = time.Now().Unix()
		info.FlushOnlineTime(false)
		a.SetUserData(info)
	}
	a.WriteMsg(&msg.C2S_UserHeartBeat{Ptime: time.Now().Unix()})
}

func handleUserLogin(args []interface{}) {
	m := args[0].(*msg.C2S_UserLogin)
	a := args[1].(gate.Agent)
	// 解析token包含的信息
	t, err := token.DecodeTokenByStr(m.Token)
	if err != nil {
		//log.Error("%v 解析token失败:%v,err:%v", m.LoginPlat, m.Token, err)
		a.Close()
		return
	} else {
		//log.Debugf("%v 平台token解析:uid:%v,device:%v,roleid:%v", m.LoginPlat, t.Uid, t.UserData, t.RoleId)
	}
	uid := uint64(stringutils.String2Int64(t.Id))
	a.SetUserData(UserInfo{Uid: uid, Ptime: 0, LoginPlat: m.LoginPlat, LoginTime: time.Now().Unix(), LastAddTime: time.Now().Unix()})
	addUserAgent(uid, a, m.LoginPlat)

	//登录返回
	a.WriteMsg(&msg.S2C_UserLogin{Ret: 1, Msg: "登录成功"})
}

func handleMessage(args []interface{}) {
	m := args[0].(*msg.C2S_Message)
	a := args[1].(gate.Agent)
	info, ok := a.UserData().(UserInfo)
	if !ok || info.Uid == 0 {
		log.Error("发送消息的用户未注册")
		return
	}
	log.Infof("接受客户端websocket消息:%v message:%+v", a.RemoteAddr(), m)
}

