package signs

import (
	"github.com/HimbeerserverDE/mt-multiserver-proxy"
)

type ClickEvent interface {
	Click(cc *proxy.ClientConn, sign *Sign)
}

type Hop struct {
	Srv string
}

func (hop *Hop) Click(cc *proxy.ClientConn, _ *Sign) {
	cc.Log("<>", "CLICK")
	cc.Hop(hop.Srv)
}
