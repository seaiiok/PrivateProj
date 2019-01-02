package net

import (
	"net/http"

	"ifix.cbpc/cbpc/api/server.api"
	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/protocol"
)

// MidController ...
func MidController(w http.ResponseWriter, p *protocol.Protocol) {

	switch p.MeConf[conf.ConstReqID] {
		
	case conf.ConstReqID1th:
		midController1th(w, p)

	case conf.ConstReqID2th:
		midController2th(w, p)

	case conf.ConstReqID3th:
		midController2th(w, p)

	case conf.ConstReqID4th:
		midController4th(w, p)

	case conf.ConstReqID5th:
		midController5th(w, p)

	case conf.ConstReqID6th:
		midController6th(w, p)

	default:

	}

}

// midController1th ...
func midController1th(w http.ResponseWriter, p *protocol.Protocol) {
	midController1thCom(w, p)
}

// midController2th ...
func midController2th(w http.ResponseWriter, p *protocol.Protocol) {
	midController2thCom(w, p)
}

// midController3th ...
func midController3th(w http.ResponseWriter, p *protocol.Protocol) {

	switch p.MeConf[conf.ConstClientsTask] {
	case conf.ConstClientsTaskD3weight:
		cli := new(api.D3Weight)
		cli.SqlSelectCols = conf.Config[conf.ConstServerD3WeightCols]
		var c api.Consumer = cli
		err := api.GetConsumerKeys(c)
		p.Data = cli.Cols

		midController3thCom(w, p, err)

	case conf.ConstClientsTaskOnlineLog:
		cli := new(api.OnlineLog)
		cli.SqlSelectCols = conf.Config[conf.ConstServerOnlineLogCols]
		var c api.Consumer = cli
		err := api.GetConsumerKeys(c)
		p.Data = cli.Cols

		midController3thCom(w, p, err)

	default:

	}
}

// midController4th ...
func midController4th(w http.ResponseWriter, p *protocol.Protocol) {

	switch p.MeConf[conf.ConstClientsTask] {
	case conf.ConstClientsTaskD3weight:
		cli := new(api.D3Weight)
		cli.SqlInsertsRows = conf.Config[conf.ConstServerD3WeightRows]
		var c api.Consumer = cli
		midController4thCom(w, p, c)

	case conf.ConstClientsTaskOnlineLog:
		cli := new(api.OnlineLog)
		cli.SqlInsertsRows = conf.Config[conf.ConstServerOnlineLogRows]
		var c api.Consumer = cli
		midController4thCom(w, p, c)

	default:

	}
}

// midController5th ...
func midController5th(w http.ResponseWriter, p *protocol.Protocol) {

	switch p.MeConf[conf.ConstClientsTask] {
	case conf.ConstClientsTaskD3weight:
	case conf.ConstClientsTaskOnlineLog:
	default:

	}
}

// midController6th ...
func midController6th(w http.ResponseWriter, p *protocol.Protocol) {

	switch p.MeConf[conf.ConstClientsTask] {
	case conf.ConstClientsTaskD3weight:
	case conf.ConstClientsTaskOnlineLog:
	default:

	}
}
