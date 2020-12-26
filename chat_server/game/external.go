package game

import (
	"gin-micro/chat_server/game/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)
