package net

import (
	"net/http"

	"ifix.cbpc/cbpc/api/server.api"
	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/convert"
	"ifix.cbpc/cbpc/pkg/protocol"
	"ifix.cbpc/cbpc/pkg/utils"
)

// MidController ...
func MidController(w http.ResponseWriter, p *protocol.Protocol) {
	p.SetProcessTrace(p.MeConf[conf.ConstReqID])
	midErrorHandle(p)
	switch p.MeConf[conf.ConstReqID] {

	case conf.ConstReqID1th:

		midController1th(w, p)

	case conf.ConstReqID2th:

		midController2th(w, p)

	case conf.ConstReqID3th:

		midController3th(w, p)

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
		res := utils.ArrayDiff(p.Data, cli.Cols)
		p.Data = make([]string, 0)
		p.Data = res
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
		cli.Rows = p.Datas
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
	p.MeConf[conf.ConstReqID] = conf.ConstReqID2th
	b, err := convert.Struct2Arraybytes(p)
	if err != nil {
		p.MeConf[conf.ConstReqID] = conf.ConstReqID6th
	}
	w.Write(b)
}

// midController6th ...
func midController6th(w http.ResponseWriter, p *protocol.Protocol) {
	p.MeConf[conf.ConstReqID] = conf.ConstReqID1th
	b, err := convert.Struct2Arraybytes(p)
	if err != nil {
		p.MeConf[conf.ConstReqID] = conf.ConstReqID6th
	}
	w.Write(b)
}
