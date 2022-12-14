package command_test

import (
	"context"
	"testing"

	"github.com/adisubagja/xcore/app/dispatcher"
	"github.com/adisubagja/xcore/app/log"
	. "github.com/adisubagja/xcore/app/log/command"
	"github.com/adisubagja/xcore/app/proxyman"
	_ "github.com/adisubagja/xcore/app/proxyman/inbound"
	_ "github.com/adisubagja/xcore/app/proxyman/outbound"
	"github.com/adisubagja/xcore/common"
	"github.com/adisubagja/xcore/common/serial"
	"github.com/adisubagja/xcore/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
