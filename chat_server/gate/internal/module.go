package internal

import (
	"gin-micro/chat_server/conf"
	"gin-micro/chat_server/game"
	"gin-micro/chat_server/msg"
	"github.com/hero1s/golib/connsvr/gate"
	"github.com/hero1s/golib/log"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Config.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Config.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		TCPAddr:         conf.Config.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		AgentChanRPC:    game.ChanRPC,
	}

	switch conf.Encoding {
	case "json":
		m.Gate.Processor = msg.JSONProcessor
	case "protobuf":
		m.Gate.Processor = msg.ProtobufProcessor
	default:
		log.Error("unknown encoding: %v", conf.Encoding)
	}
}
