package app

import (
	"fmt"

	"ifix.cbpc/cbpc/interfaces"
	"ifix.cbpc/cbpc/internal/pkg/net"
	"ifix.cbpc/cbpc/internal/pkg/service"
	 "ifix.cbpc/cbpc/pkg/conf" //get user config
)

// ServiceStart ...
func ServiceStart() {
	fmt.Println("http client start")
	var o interfaces.Object
	p := new(service.Objects)
	p.DeviceTask=conf.ConstTaskD3Weight
	p.DeviceRouter= "routerSetObjectInfo"
	o = p
	status := net.NetPost(&o)
	fmt.Println(status)
	fmt.Println(p)
	status = net.NetPost(&o)
	fmt.Println(status)
	fmt.Println(p)
}
