package core_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/adisubagja/xcore/app/dispatcher"
	"github.com/adisubagja/xcore/app/proxyman"
	"github.com/adisubagja/xcore/common"
	"github.com/adisubagja/xcore/common/net"
	"github.com/adisubagja/xcore/common/protocol"
	"github.com/adisubagja/xcore/common/serial"
	"github.com/adisubagja/xcore/common/uuid"
	. "github.com/adisubagja/xcore/core"
	"github.com/adisubagja/xcore/features/dns"
	"github.com/adisubagja/xcore/features/dns/localdns"
	_ "github.com/adisubagja/xcore/main/distro/all"
	"github.com/adisubagja/xcore/proxy/dokodemo"
	"github.com/adisubagja/xcore/proxy/vmess"
	"github.com/adisubagja/xcore/proxy/vmess/outbound"
	"github.com/adisubagja/xcore/testing/servers/tcp"
)

func TestXrayDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	})
	instance.AddFeature(localdns.New())
	<-wait
}

func TestXrayClose(t *testing.T) {
	port := tcp.PickPort()

	userID := uuid.New()
	config := &Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{
						Range: []*net.PortRange{net.SinglePortRange(port)},
					},
					Listen: net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userID.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}
