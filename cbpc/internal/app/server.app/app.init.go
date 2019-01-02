package app

import (
	"ifix.cbpc/cbpc/internal/pkg/net.server"
	_ "ifix.cbpc/cbpc/pkg/conf" //get user config
)

func init() {

}

func ServiceStart() {
	net.ServerHTTPStart()
}
