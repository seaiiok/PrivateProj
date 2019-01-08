package net

import (
	"fmt"
	"net/http"

	"ifix.cbpc/cbpc/interfaces"
	"ifix.cbpc/cbpc/internal/pkg/service"
	"ifix.cbpc/cbpc/pkg/convert"
)

func routerSetObjectInfo(w http.ResponseWriter, r *http.Request) {
	var o interfaces.Object
	p := new(service.Objects)
	o = p
	err := convert.Reader2Struct(r.Body, o)
	p.SetError(err)

	o.GetObjectInfo()
	p.SetRouter("routerSetObjectConf")
	b, _ := convert.Struct2Arraybytes(o)
	w.Write(b)
}

func routerSetObjectConf(w http.ResponseWriter, r *http.Request) {
	var o interfaces.Object
	p := new(service.Objects)
	o = p
	err := convert.Reader2Struct(r.Body, o)
	p.SetError(err)

	o.GetObjectConf()
	p.SetRouter("routerSetObjectKeys")
	b, _ := convert.Struct2Arraybytes(o)
	w.Write(b)
}

func routerSetObjectKeys(w http.ResponseWriter, r *http.Request) {
	fmt.Println("routerSetObjectKeys")
}

func routerSetObjectDatas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("routerSetObjectDatas")
}

func routerSetObjectRestart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("routerSetObjectRestart")
}
