package multiaddr

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

const P_RNDZ = 0x300000

func init() {
	protoRNDZ := ma.Protocol{
		Name:  "rndz",
		Code:  P_RNDZ,
		Size:  0,
		VCode: ma.CodeToVarint(P_RNDZ),
	}

	if err := ma.AddProtocol(protoRNDZ); err != nil {
		panic(err)
	}
}

func NewListenAddr(localAddr, rndzServer ma.Multiaddr) ma.Multiaddr {
	return rndzServer.Encapsulate(rndzServer)
}

func SplitListenAddr(addr ma.Multiaddr) (localAddr ma.Multiaddr, rndzAddr ma.Multiaddr) {
	localAddr, rndzAddr = ma.SplitFunc(addr, func(c ma.Component) bool {
		return c.Protocol().Code == P_RNDZ
	})

	if rndzAddr != nil {
		_, rndzAddr = ma.SplitFirst(rndzAddr)
	}

	return
}

func NewDialAddr(rndzServer ma.Multiaddr, peerId peer.ID) ma.Multiaddr {
	p2pPart, err := ma.NewMultiaddr(fmt.Sprintf("/p2p/%s", peerId))
	if err != nil {
		panic(err)
	}

	return rndzServer.Encapsulate(p2pPart)
}

func SplitDialAddr(addr ma.Multiaddr) (ma.Multiaddr, peer.ID) {
	return peer.SplitAddr(addr)
}
