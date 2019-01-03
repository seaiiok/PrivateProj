package net

import (
	"errors"
	"time"

	"ifix.cbpc/cbpc/api/clients.api"
	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/protocol"
)

// MidController ...
func MidController(p *protocol.Protocol) {
	switch p.MeConf[conf.ConstReqID] {

	case conf.ConstReqID1th:
		midController1th(p)

	case conf.ConstReqID2th:
		midController2th(p)

	case conf.ConstReqID3th:
		midController3th(p)

	case conf.ConstReqID4th:
		midController4th(p)

	case conf.ConstReqID5th:
		midController5th(p)

	case conf.ConstReqID6th:
		midController6th(p)

	default:

	}

}

// midController1th ...
func midController1th(p *protocol.Protocol) {
	if p.Error != nil {
		clientconf := make(map[string]string, 0)
		clientconf[conf.ConstServerAddr] = p.MeConf[conf.ConstServerAddr]
		clientconf[conf.ConstServerPort] = p.MeConf[conf.ConstServerPort]
		clientconf[conf.ConstClientsTask] = p.MeConf[conf.ConstClientsTask]
		p.Error = conf.SetConf("ifixtools.conf", clientconf)
	}

}

// midController2th ...
func midController2th(p *protocol.Protocol) {
	switch p.MeConf[conf.ConstClientsTask] {
	case conf.ConstClientsTaskD3weight:
		cli := new(api.Empty)
		cli.DriverName = p.MeConf[conf.ConstClientsD3WeightDriverName]
		cli.DataSourceName = p.MeConf[conf.ConstClientsD3WeightDataSourceName]
		var c api.Producer = cli
		p.Error = api.GetProducerPing(c)

	case conf.ConstClientsTaskOnlineLog:
	default:
	}
}

// midController3th ...
func midController3th(p *protocol.Protocol) {

	switch p.MeConf[conf.ConstClientsTask] {
	case conf.ConstClientsTaskD3weight:
		cli := new(api.D3Weight)
		cli.SqlSelectCols = p.MeConf[conf.ConstClientsD3WeightCols]
		var c api.Producer = cli
		p.Error = api.GetProducerKeys(c)
		p.Data = cli.Cols

	case conf.ConstClientsTaskOnlineLog:

	default:

	}

}

// midController4th ...
func midController4th(p *protocol.Protocol) {

	switch p.MeConf[conf.ConstClientsTask] {
	case conf.ConstClientsTaskD3weight:
		cli := new(api.D3Weight)

		cli.SqlSelectRows = midController4thCom(p)
		if cli.SqlSelectRows == "" {
			p.Error = errors.New("nothing datas")
			return
		}
		var c api.Producer = cli
		p.Error = api.GetProducerRows(c)
		p.Datas = cli.Rows
	case conf.ConstClientsTaskOnlineLog:

	default:

	}
}

// midController5th ...
func midController5th(p *protocol.Protocol) {
	time.Sleep(conf.ConstWaitTime * time.Second)
}

// midController6th ...
func midController6th(p *protocol.Protocol) {
	time.Sleep(conf.ConstWaitTime * time.Second)

}
