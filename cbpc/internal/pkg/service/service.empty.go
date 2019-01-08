package service

import (
	"ifix.cbpc/cbpc/proto"
)

//Objects ...
type Objects struct {
	proto.Proto
}

//GetObjectInfo ...
func (p *Objects) GetObjectInfo() {
	p.getObjectInfo()
}

//GetObjectConf ...
func (p *Objects) GetObjectConf() {
	p.getObjectConf()
}

//GetObjectKeys ...
func (p *Objects) GetObjectKeys() {
}

//GetObjectDatas ...
func (p *Objects) GetObjectDatas() {
}

//SetObjectInfo ...
func (p *Objects) SetObjectInfo() {
}

//SetObjectConf ...
func (p *Objects) SetObjectConf() {
}

//SetObjectKeys ...
func (p *Objects) SetObjectKeys() {
}

//SetObjectDatas ...
func (p *Objects) SetObjectDatas() {
}
