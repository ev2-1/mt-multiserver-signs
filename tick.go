package signs

import (
	"github.com/ev2-1/mt-multiserver-playertools"

	"sync"
)

var initUpdatesMu sync.Once

func initUpdates() {
	initUpdatesMu.Do(func() {
		playerTools.RegisterPlayerListUpdateHandler(&playerTools.PlayerListUpdateHandler{
			Update: func(_ playerTools.PlayerList) {
				Update()
			},
		})
		playerTools.RegisterSrvPlayerListHandler(&playerTools.SrvPlayerListHandler{
			Update: func(_ playerTools.PlayerList, server string) {
				Update()
			},
		})
	})
}
