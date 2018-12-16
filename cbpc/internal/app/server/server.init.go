package server

import (
	"ifix.cbpc/cbpc/internal/pkg"
)

func serverInit() {
	pkg.ServerConfigInit("")
	pkg.ServerDBInit()
}
