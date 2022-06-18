package signs

import (
	"github.com/anon55555/mt"
	
	"github.com/HimbeerserverDE/mt-multiserver-proxy"
)

// TODO: readyplayers correctly (has to be implemented in multiserver itself)
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
