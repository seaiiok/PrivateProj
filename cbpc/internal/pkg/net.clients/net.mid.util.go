package net

import (
	"net/http"

	"ifix.cbpc/cbpc/api/server.api"
	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/convert"
	"ifix.cbpc/cbpc/pkg/protocol"
)

func midController1thCom(w http.ResponseWriter, p *protocol.Protocol) {
	for k, v := range conf.Config {
		p.MeConf[k] = v
	}

	p.MeConf[conf.ConstReqID] = conf.ConstReqID2th
	b, err := convert.Struct2Arraybytes(p)
	if err != nil {
		p.MeConf[conf.ConstReqID] = conf.ConstReqID6th
	}
	w.Write(b)
}

func midController2thCom(w http.ResponseWriter, p *protocol.Protocol) {
	cli := new(api.Empty)
	cli.DriverName = conf.Config[conf.ConstServerDriverName]
	cli.DataSourceName = conf.Config[conf.ConstServerDataSourceName]
	var c api.Consumer = cli
	err := api.GetConsumerPing(c)

	p.MeConf[conf.ConstReqID] = conf.ConstReqID3th
	b, err := convert.Struct2Arraybytes(p)
	if err != nil {
		p.MeConf[conf.ConstReqID] = conf.ConstReqID6th
	}
	w.Write(b)
}

func midController3thCom(w http.ResponseWriter, p *protocol.Protocol, err error) {
	p.MeConf[conf.ConstReqID] = conf.ConstReqID4th
	b, err := convert.Struct2Arraybytes(p)
	if err != nil {
		p.MeConf[conf.ConstReqID] = conf.ConstReqID6th
	}
	w.Write(b)
}

func midController4thCom(w http.ResponseWriter, p *protocol.Protocol, c api.Consumer) {
	err := api.SetConsumerRows(c)
	p.MeConf[conf.ConstReqID] = conf.ConstReqID5th
	b, err := convert.Struct2Arraybytes(p)
	if err != nil {
		p.MeConf[conf.ConstReqID] = conf.ConstReqID6th
	}
	w.Write(b)
}
