package net

import (
	"fmt"
	"net/http"

	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/protocol"
)

var client *http.Client

func ClientsHTTPStart() {
	ClientsHTTPInit()
	p := protocol.NewProto()
	GetMeTask(p)
	
	p.MeConf[conf.ConstReqID] = conf.ConstReqID1th

	for {
		p = HTTPPost(p)
		MidController(p)
		fmt.Println(p.ProcessTrace)
	}

}
