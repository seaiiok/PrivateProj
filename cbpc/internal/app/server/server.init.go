package server

import (
	"fmt"

	"ifix.cbpc/cbpc/internal/pkg"

	//	"ifix.cbpc/cbpc/internal/pkg"
)

func Start() {
	var d3 pkg.D3Weight
	var p3 pkg.Proto
	var Intf pkg.IServerDatebase
	Intf = &d3
	pkg.ConfigInit("")
	p := Intf.ServerDatebaseGetReady(&p3)
	fmt.Println(p)
}
