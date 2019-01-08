package service

import (
	"ifix.cbpc/cbpc/pkg/conf"
)

func (p *Objects) getObjectConf() {
	switch p.DeviceTask {
	case conf.ConstTaskD3Weight:
		p.Query=conf.Config[conf.ConstD3WeightServerKeys]
		p.GetFulQuery()
		p.Insert=conf.Config[conf.ConstD3WeightServerData]
	case conf.ConstTaskOnlineLog:

	case conf.ConstTaskOfflineZip:

	default:

	}
}
