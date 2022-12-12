package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/adisubagja/xcore/app/dispatcher"
	_ "github.com/adisubagja/xcore/app/proxyman/inbound"
	_ "github.com/adisubagja/xcore/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/adisubagja/xcore/app/commander"
	_ "github.com/adisubagja/xcore/app/log/command"
	_ "github.com/adisubagja/xcore/app/proxyman/command"
	_ "github.com/adisubagja/xcore/app/stats/command"

	// Developer preview services
	_ "github.com/adisubagja/xcore/app/observatory/command"

	// Other optional features.
	_ "github.com/adisubagja/xcore/app/dns"
	_ "github.com/adisubagja/xcore/app/dns/fakedns"
	_ "github.com/adisubagja/xcore/app/log"
	_ "github.com/adisubagja/xcore/app/metrics"
	_ "github.com/adisubagja/xcore/app/policy"
	_ "github.com/adisubagja/xcore/app/reverse"
	_ "github.com/adisubagja/xcore/app/router"
	_ "github.com/adisubagja/xcore/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/adisubagja/xcore/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/adisubagja/xcore/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/adisubagja/xcore/proxy/blackhole"
	_ "github.com/adisubagja/xcore/proxy/dns"
	_ "github.com/adisubagja/xcore/proxy/dokodemo"
	_ "github.com/adisubagja/xcore/proxy/freedom"
	_ "github.com/adisubagja/xcore/proxy/http"
	_ "github.com/adisubagja/xcore/proxy/loopback"
	_ "github.com/adisubagja/xcore/proxy/mtproto"
	_ "github.com/adisubagja/xcore/proxy/shadowsocks"
	_ "github.com/adisubagja/xcore/proxy/socks"
	_ "github.com/adisubagja/xcore/proxy/trojan"
	_ "github.com/adisubagja/xcore/proxy/vless/inbound"
	_ "github.com/adisubagja/xcore/proxy/vless/outbound"
	_ "github.com/adisubagja/xcore/proxy/vmess/inbound"
	_ "github.com/adisubagja/xcore/proxy/vmess/outbound"
	_ "github.com/adisubagja/xcore/proxy/wireguard"

	// Transports
	_ "github.com/adisubagja/xcore/transport/internet/domainsocket"
	_ "github.com/adisubagja/xcore/transport/internet/grpc"
	_ "github.com/adisubagja/xcore/transport/internet/http"
	_ "github.com/adisubagja/xcore/transport/internet/kcp"
	_ "github.com/adisubagja/xcore/transport/internet/quic"
	_ "github.com/adisubagja/xcore/transport/internet/tcp"
	_ "github.com/adisubagja/xcore/transport/internet/tls"
	_ "github.com/adisubagja/xcore/transport/internet/udp"
	_ "github.com/adisubagja/xcore/transport/internet/websocket"
	_ "github.com/adisubagja/xcore/transport/internet/xtls"

	// Transport headers
	_ "github.com/adisubagja/xcore/transport/internet/headers/http"
	_ "github.com/adisubagja/xcore/transport/internet/headers/noop"
	_ "github.com/adisubagja/xcore/transport/internet/headers/srtp"
	_ "github.com/adisubagja/xcore/transport/internet/headers/tls"
	_ "github.com/adisubagja/xcore/transport/internet/headers/utp"
	_ "github.com/adisubagja/xcore/transport/internet/headers/wechat"
	_ "github.com/adisubagja/xcore/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/adisubagja/xcore/main/json"
	_ "github.com/adisubagja/xcore/main/toml"
	_ "github.com/adisubagja/xcore/main/yaml"

	// Load config from file or http(s)
	_ "github.com/adisubagja/xcore/main/confloader/external"

	// Commands
	_ "github.com/adisubagja/xcore/main/commands/all"
)
