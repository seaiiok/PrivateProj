package producers

import (
	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/db"
	"ifix.cbpc/cbpc/pkg/protocol"
	"ifix.cbpc/cbpc/pkg/utils"
)

//D3weight ...
type D3weight struct {
	protocol.Proto
}

//SetObjectInfo ...
func (p *D3weight) SetObjectInfo() {
	p.SetErrorNil()
	p.Device.DeviceIP = utils.GetLocalMatchIp(conf.ConstMatchIP)
	p.Device.DeviceName = "3#门车辆称重系统"
	p.Device.DeviceTask = conf.Config[conf.ConstClientsTask]
	p.Device.DeviceRouter = conf.ConstRouterGetObjectInfo
	p.SQL.DatabaseDriver = conf.Config[conf.ConstClientsDriver]
	p.SQL.DatabaseSource = conf.Config[conf.ConstClientsSource]
}

//SetObjectConf ...
func (p *D3weight) SetObjectConf() {
	err := db.Init(p.SQL.DatabaseDriver, p.SQL.DatabaseSource)
	if err != nil {
		p.SetError(true)
	}
	p.SQL.DatabaseDriver = conf.Config[conf.ConstServerDriver]
	p.SQL.DatabaseSource = conf.Config[conf.ConstServerSource]
}

//SetObjectKeys ...
func (p *D3weight) SetObjectKeys() {
	var err error
	p.SQL.QuerySQL = conf.Config[conf.ConstClientsKeys]
	p.SQL.Args, err = db.Query(p.SQL.QuerySQL)
	if err != nil {
		p.SetError(true)
	}
	p.SQL.QuerySQL = conf.Config[conf.ConstServerKeys]
}

//SetObjectDatas ...
func (p *D3weight) SetObjectDatas() {
	var err error
	p.SQL.QuerySQL = conf.Config[conf.ConstClientsData]
	p.SQL.QuerySQL, err = getFulQuery(p.SQL.QuerySQL, p.SQL.Args)
	if err != nil {
		p.SetError(true)
	} else {
		p.SQL.Data, err = db.Querys(p.SQL.QuerySQL)
	}
	if err != nil {
		p.SetError(true)
	}
	p.SQL.InsertSQL = conf.Config[conf.ConstServerData]
}
