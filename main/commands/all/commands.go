package all

import (
	"github.com/adisubagja/xcore/main/commands/all/api"
	"github.com/adisubagja/xcore/main/commands/all/tls"
	"github.com/adisubagja/xcore/main/commands/base"
)

// go:generate go run github.com/adisubagja/xcore/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
	)
}
