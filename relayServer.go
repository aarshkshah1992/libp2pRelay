package main

import (
	"context"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
	quic "github.com/libp2p/go-libp2p-quic-transport"
	"github.com/libp2p/go-tcp-transport"
	ma "github.com/multiformats/go-multiaddr"
)

func main() {
	publicIP := os.Getenv("RELAY_IP")
	factory := func(addrs []ma.Multiaddr) []ma.Multiaddr {
		if len(publicIP) != 0 {
			tcp := fmt.Sprintf("/ip4/%s/tcp/12001", publicIP)
			quic := fmt.Sprintf("/ip4/%s/udp/12001/quic", publicIP)
			return append(addrs, ma.StringCast(tcp), ma.StringCast(quic))
		}

		return addrs
	}

	// A public relay server that supports TCP & QUIC and listens on port 12001
	ctx := context.Background()
	h1, err := libp2p.New(ctx, libp2p.ForceReachabilityPublic(), libp2p.EnableRelay(circuit.OptHop),
		libp2p.AddrsFactory(factory),
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(quic.NewTransport),
		libp2p.ListenAddrs(ma.StringCast("/ip4/0.0.0.0/tcp/12001"),
			ma.StringCast("/ip4/0.0.0.0/udp/12001/quic")))
	if err != nil {
		panic(err)
	}

	fmt.Println("\n relay server peerID: ", h1.ID().Pretty())
	fmt.Println("\n relay server addresses:")
	for _, a := range h1.Addrs() {
		fmt.Println(a)
	}

	// Relay connections
	for {

	}
}
