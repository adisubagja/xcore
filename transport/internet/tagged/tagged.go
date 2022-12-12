package tagged

import (
	"context"

	"github.com/adisubagja/xcore/common/net"
)

type DialFunc func(ctx context.Context, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
