package app

import (
	"fmt"

	"ifix.cbpc/cbpc/internal/pkg/net.clients"
	_ "ifix.cbpc/cbpc/pkg/conf" //get user config
)

func init() {

}

func ServiceStart() {
	fmt.Println("http client start")
	net.ClientsHTTPInit()
	p := new(net.Clientsproto)
	p.HTTPPost()
	fmt.Println(p)
}
