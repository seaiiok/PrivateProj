package app

import (
	"fmt"

	"ifix.cbpc/cbpc/internal/pkg/net"
)

// ServiceStart ...
func ServiceStart() {
	fmt.Println("http server start")
	net.ServerHTTPStart()
}
