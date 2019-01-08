package service

import (
	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/db"
	"ifix.cbpc/cbpc/proto"
)

//D3Weight ...
type D3Weight struct {
	proto.Proto
}

//GetObjectInfo ...
func (p *D3Weight) GetObjectInfo() {

}

//GetObjectConf ...
func (p *D3Weight) GetObjectConf() {
	config := p.Config.Config
	p.Sql.DatabaseDriver = config[conf.ConstD3WeightClientsDriver]
	p.Sql.DatabaseSource = config[conf.ConstD3WeightClientsSource]
	p.SetQuery(config[conf.ConstD3WeightClientsDriver])
	p.Err.Err = db.Init(p.Sql.DatabaseDriver, p.Sql.DatabaseSource)
}

//GetObjectKeys ...
func (p *D3Weight) GetObjectKeys() {
	config := p.Config.Config
	p.SetQuery(config[conf.ConstD3WeightServerKeys])
}

//GetObjectDatas ...
func (p *D3Weight) GetObjectDatas() {
	config := p.Config.Config
	p.SetQuery(config[conf.ConstD3WeightServerKeys])
}

//SetObjectInfo ...
func (p *D3Weight) SetObjectInfo() {

}

//SetObjectConf ...
func (p *D3Weight) SetObjectConf() {
	config := p.Config.Config
	p.Sql.DatabaseDriver = config[conf.ConstD3WeightClientsDriver]
	p.Sql.DatabaseSource = config[conf.ConstD3WeightClientsSource]
	p.SetQuery(config[conf.ConstD3WeightClientsDriver])
	p.Err.Err = db.Init(p.Sql.DatabaseDriver, p.Sql.DatabaseSource)
}

//SetObjectKeys ...
func (p *D3Weight) SetObjectKeys() {
	config := p.Config.Config
	p.SetQuery(config[conf.ConstD3WeightServerKeys])
}

//SetObjectDatas ...
func (p *D3Weight) SetObjectDatas() {
	config := p.Config.Config
	p.SetQuery(config[conf.ConstD3WeightServerKeys])
}
