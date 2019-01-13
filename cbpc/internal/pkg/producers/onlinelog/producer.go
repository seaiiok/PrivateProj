package producer

import (
	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/db"
	"ifix.cbpc/cbpc/pkg/protocol"
	"ifix.cbpc/cbpc/pkg/utils"
)

//OnlineLog ...
type OnlineLog struct {
	protocol.Proto
}

//SetObjectInfo ...
func (p *OnlineLog) SetObjectInfo() {
	p.SetErrorNil()
	p.Device.DeviceIP = utils.GetLocalMatchIp(conf.ConstMatchIP)
	p.Device.DeviceName = "在线清数日志"
	p.Device.DeviceTask = conf.Config[conf.ConstClientsTask]
	p.Device.DeviceRouter = conf.ConstRouterGetObjectInfo
	p.SQL.DatabaseDriver = conf.Config[conf.ConstClientsDriver]
	p.SQL.DatabaseSource = conf.Config[conf.ConstClientsSource]
}

//SetObjectRoute ...
func (p *OnlineLog) SetObjectRoute() string {
	return p.Device.DeviceRouter
}

//SetObjectConf ...
func (p *OnlineLog) SetObjectConf() {
	err := db.Init(p.SQL.DatabaseDriver, p.SQL.DatabaseSource)
	if err != nil {
		p.SetError(true)
	}
	p.SQL.DatabaseDriver = conf.Config[conf.ConstServerDriver]
	p.SQL.DatabaseSource = conf.Config[conf.ConstServerSource]
}

//SetObjectKeys ...
func (p *OnlineLog) SetObjectKeys() {
	var err error
	p.SQL.QuerySQL = conf.Config[conf.ConstClientsKeys]
	p.SQL.Args, err = db.Query(p.SQL.QuerySQL)
	if err != nil {
		p.SetError(true)
	}
	p.SQL.QuerySQL = conf.Config[conf.ConstServerKeys]
}

//SetObjectDatas ...
func (p *OnlineLog) SetObjectDatas() {

}
