package service

import (
	"ifix.cbpc/cbpc/pkg/conf"
)

func (p *Objects) getObjectInfo() {
	switch p.DeviceTask {
	case conf.ConstTaskD3Weight:
		p.DatabaseDriver=conf.Config[conf.ConstD3WeightClientsDriver]
		p.DatabaseSource=conf.Config[conf.ConstD3WeightClientsSource]
	case conf.ConstTaskOnlineLog:

	case conf.ConstTaskOfflineZip:

	default:

	}
}
