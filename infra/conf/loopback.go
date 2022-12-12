package conf

import (
	"github.com/golang/protobuf/proto"
	"github.com/adisubagja/xcore/proxy/loopback"
)

type LoopbackConfig struct {
	InboundTag string `json:"inboundTag"`
}

func (l LoopbackConfig) Build() (proto.Message, error) {
	return &loopback.Config{InboundTag: l.InboundTag}, nil
}
