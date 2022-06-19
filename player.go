package signs

import (
	"github.com/HimbeerserverDE/mt-multiserver-proxy"
	"github.com/anon55555/mt"

	"sync"
)

// Ready is called internaly
// Should not have to be called externaly
func Ready(cc *proxy.ClientConn) {
	updateSignText()

	// initialize client:
	add := []mt.AOAdd{}

	signsMu.RLock()
	signsMu.RUnlock()

	for _, s := range signs[cc.ServerName()] {
		add = append(add, GenerateSignAOAdd(s.cachedText, s.Color, s.Pos.Pos, s.Pos.Rotation, s.Pos.Wall, s.aoid))
	}

	if len(add) != 0 {
		cc.SendCmd(&mt.ToCltAORmAdd{
			Add: add,
		})
	}

	// client ready
	readyClients[cc.Name()] = true
}

// Unready is called internaly
// Should not have to be called externaly
func Unready(cc *proxy.ClientConn, srv string) {
	signsMu.RLock()
	defer signsMu.RUnlock()

	var AOIDs []mt.AOID

	for _, s := range signs[srv] {
		AOIDs = append(AOIDs, s.aoid)
	}

	if len(AOIDs) != 0 {
		cc.SendCmd(&mt.ToCltAORmAdd{
			Remove: AOIDs,
		})
	}
}

var initPlayerActivatorMu sync.Once

// registers all the stuffs
func initPlayerActivator() {
	initPlayerActivatorMu.Do(func() {
		proxy.RegisterClientHandler(&proxy.ClientHandler{
			AOReady: func(cc *proxy.ClientConn) {
				Ready(cc)
			},
			Hop: func(cc *proxy.ClientConn, src, dest string) {
				Unready(cc, src)
				Ready(cc)
			},
			Leave: func(cc *proxy.ClientConn, _ *proxy.Leave) {
				Unready(cc, cc.ServerName())
			},
		})
	})
}
