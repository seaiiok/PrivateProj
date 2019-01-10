package app

import (
	"net/http"

	"ifix.cbpc/cbpc/internal/pkg/consumers"
	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/convert"
	"ifix.cbpc/cbpc/pkg/interfaces"
)

//clients info
func routerGetObjectInfo(w http.ResponseWriter, r *http.Request) {
	
	var o interfaces.Consumers
	p := new(consumers.Objects)
	o = p
	err := convert.Reader2Struct(r.Body, o)
	if err != nil {
		p.SetError(true)
	}
	
	if p.Err.Error != false {
		p.SetRouter(conf.ConstRouterGetObjectConf)
		p.SetProcessTrace(p.Device.DeviceRouter)
		b, _ := convert.Struct2Arraybytes(o)
		w.Write(b)

		return
	}

	p.SetRouter(conf.ConstRouterGetObjectConf)
	p.SetProcessTrace(p.Device.DeviceRouter)
	b, _ := convert.Struct2Arraybytes(o)
	w.Write(b)
}

//server config
func routerGetObjectConf(w http.ResponseWriter, r *http.Request) {
	var o interfaces.Consumers
	p := new(consumers.Objects)
	o = p
	err := convert.Reader2Struct(r.Body, o)
	if err != nil {
		p.SetError(true)
	}
	o.GetObjectConf()

	if p.Err.Error != false {
		p.SetRouter(conf.ConstRouterGetObjectKeys)
		p.SetProcessTrace(p.Device.DeviceRouter)
		b, _ := convert.Struct2Arraybytes(o)
		w.Write(b)

		return
	}
	p.SetRouter(conf.ConstRouterGetObjectKeys)
	p.SetProcessTrace(p.Device.DeviceRouter)
	b, _ := convert.Struct2Arraybytes(o)
	w.Write(b)
}

//get server keys
func routerGetObjectKeys(w http.ResponseWriter, r *http.Request) {
	var o interfaces.Consumers
	p := new(consumers.Objects)
	o = p
	err := convert.Reader2Struct(r.Body, o)
	if err != nil {
		p.SetError(true)
	}

	if p.Err.Error != false {
		p.SetRouter(conf.ConstRouterGetObjectDatas)
		p.SetProcessTrace(p.Device.DeviceRouter)
		b, _ := convert.Struct2Arraybytes(o)
		w.Write(b)
		return
	}
	o.GetObjectKeys()
	p.SetRouter(conf.ConstRouterGetObjectDatas)
	p.SetProcessTrace(p.Device.DeviceRouter)
	b, _ := convert.Struct2Arraybytes(o)
	w.Write(b)
}

//get clients data
func routerGetObjectDatas(w http.ResponseWriter, r *http.Request) {
	var o interfaces.Consumers
	p := new(consumers.Objects)
	o = p
	err := convert.Reader2Struct(r.Body, o)
	if err != nil {
		p.SetError(true)
	}

	if p.Err.Error != false {
		p.SetRouter(conf.ConstRouterGetObjectAgain)
		p.SetProcessTrace(p.Device.DeviceRouter)
		b, _ := convert.Struct2Arraybytes(o)
		w.Write(b)
		return
	}
	o.GetObjectDatas()
	p.SetRouter(conf.ConstRouterGetObjectAgain)
	p.SetProcessTrace(p.Device.DeviceRouter)
	b, _ := convert.Struct2Arraybytes(o)
	w.Write(b)
}

//restart service
func routerGetObjectRestart(w http.ResponseWriter, r *http.Request) {
	var o interfaces.Consumers
	p := new(consumers.Objects)
	o = p

	err := convert.Reader2Struct(r.Body, o)
	if err != nil {
		p.SetError(true)
	}
	p.SetRouter(conf.ConstRouterGetObjectInfo)
	p.SetProcessTrace(p.Device.DeviceRouter)
	b, _ := convert.Struct2Arraybytes(o)
	w.Write(b)
}

//again service
func routerGetObjectAgain(w http.ResponseWriter, r *http.Request) {
	var o interfaces.Consumers
	p := new(consumers.Objects)
	o = p
	err := convert.Reader2Struct(r.Body, o)
	if err != nil {
		p.SetError(true)
	}
	p.SetRouter(conf.ConstRouterGetObjectConf)
	p.SetProcessTrace(p.Device.DeviceRouter)
	b, _ := convert.Struct2Arraybytes(o)
	w.Write(b)
}
