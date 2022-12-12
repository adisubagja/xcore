package tcp

import (
	"github.com/adisubagja/xcore/common"
	"github.com/adisubagja/xcore/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
