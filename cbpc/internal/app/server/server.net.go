package server

import (
	// "fmt"
	"ifix.cbpc/cbpc/internal/pkg"
)

//ServiceStart app main service
func ServiceStart() {
	serverInit()
	pkg.ServerHTTPStart()
}