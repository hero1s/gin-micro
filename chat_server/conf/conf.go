package conf

import (
	"time"
)

var (
	// gate conf
	Encoding               = "json" // "json" or "protobuf"
	PendingWriteNum        = 2000
	MaxMsgLen       uint32 = 8192
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 2
	LittleEndian           = false

	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	ChanRPCLen         = 10000
)
