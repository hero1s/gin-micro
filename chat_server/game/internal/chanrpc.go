package internal

import (
	"gin-micro/chat_server/msg"
	"github.com/hero1s/golib/connsvr/gate"
	"github.com/hero1s/golib/log"
	"github.com/hero1s/golib/utils/util"
	"sync"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

var users util.Map
var mapUsersByUid util.Map
var mutex sync.Mutex

func getUserAgent(uid uint64) gate.Agent {
	agent := mapUsersByUid.Get(uid)
	if agent != nil {
		return agent.(gate.Agent)
	}
	return nil
}

func getAgentBindUid(a gate.Agent) uint64 {
	info, ok := a.UserData().(UserInfo)
	if !ok {
		return 0
	}
	return info.Uid
}

func addUserAgent(uid uint64, a gate.Agent, loginPlat int) {
	mutex.Lock()
	defer mutex.Unlock()
	oldAgent := getUserAgent(uid)
	if oldAgent != nil && a != oldAgent {
		log.Error("平台:%v 新注册socket顶掉旧的socket:%v--%v,%v", loginPlat, uid, a.RemoteAddr(), oldAgent.RemoteAddr())
		oldAgent.SetUserData(UserInfo{Uid: 0, Ptime: 0, LoginPlat: loginPlat})
		oldAgent.Close() //toney 可能有Bug
		mapUsersByUid.Del(uid)
		users.Del(oldAgent)
	}
	mapUsersByUid.Set(uid, a)
	//log.Infof("注册 %v 平台用户:%v--%v,记名用户:%v,总用户:%v", loginPlat, uid, a.RemoteAddr(), len(mapUsersByUid), len(users))
}

func rpcNewAgent(args []interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	a := args[0].(gate.Agent)
	users.Set(a, struct{}{})
	log.Debugf("客户端新连接:ip:%v,在线人数:%v", a.RemoteAddr(), users.Len())
	a.SetUserData(UserInfo{Uid: 0, Ptime: 0, LoginPlat: 0, LoginTime: 0})
	a.WriteMsg(&msg.S2C_Message{Uid: 0, Message: "hello,come on"})
}

func rpcCloseAgent(args []interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	a := args[0].(gate.Agent)
	users.Del(a)
	info, ok := a.UserData().(UserInfo)
	if !ok || info.Uid == 0 {
		log.Debugf("未登录客户端断开连接:%v,在线人数:%v", a.RemoteAddr(), users.Len())
		return
	}
	info.FlushOnlineTime(true)
	mapUsersByUid.Del(info.Uid)
	log.Infof("移除用户:%v--%v,%v平台记名用户:%v,总用户:%v", info.Uid, a.RemoteAddr(), info.LoginPlat, mapUsersByUid.Len(), users.Len())
}

func broadcastMsg(msg interface{}, _a gate.Agent) {
	mutex.Lock()
	defer mutex.Unlock()
	log.Debugf("websocket广播消息:%+v,当前用户数:%v--%v", msg, mapUsersByUid.Len(), users.Len())
	users.LockRange(func(k interface{}, v interface{}) {
		if k != _a && k != nil {
			k.(gate.Agent).WriteMsg(msg)
		}
	})
}
func sendMessage(msg interface{}, uids []uint64, isLog bool) {
	mutex.Lock()
	defer mutex.Unlock()
	for _, uid := range uids {
		a := getUserAgent(uid)
		if a != nil {
			a.WriteMsg(msg)
			if isLog {
				log.Infof("给用户 %v 发送消息:%+v", uid, msg)
			}
		} else {
			log.Debugf("发现消息用户网络未注册:%v,消息:%+v", uid, msg)
		}
	}
}


