package main

import (
	//"strings"
	proxy "github.com/HimbeerserverDE/mt-multiserver-proxy"
	"github.com/anon55555/mt"
)

var signs = []string{"mcl_signs:wall_sign","mcl_signs:standing_sign22_5","mcl_signs:standing_sign45","mcl_signs:standing_sign67_5"}

func log(cc *proxy.ClientConn, cmd *mt.ToSrvInteract) bool {
	cc.Log("<-", "interaction", cmd.Action.String())

	switch pos := cmd.Pointed.(type) {
	case *mt.PointedNode:
		cc.Log(" ^", "Node", pos.Under[0], pos.Under[1], pos.Under[2])
	}

	return false
}

func init() {
	for _, sign := range signs {
		proxy.RegisterNodeHandler(&proxy.NodeHandler{
			Node: sign,
			OnDig: log,
		})
	}
}
