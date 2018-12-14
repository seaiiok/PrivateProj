package server

import (
	"fmt"
	"ifix.cbpc/cbpc/internal/pkg"
)

//ServiceStart app main service
func ServiceStart() {
	serverInit()
	// var d3 pkg.D3Weight
	// var p3 pkg.Proto
	// var iServer pkg.IServer
	// iServer = &d3
	// p := iServer.ServerDatebaseGetReady(&p3)
	// fmt.Println(p)
	pkg.ServerHTTPStart()
}