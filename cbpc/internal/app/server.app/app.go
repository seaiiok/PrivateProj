package app

import (
	"ifix.cbpc/cbpc/pkg/conf"
	"net/http"
)
func ServiceStart(){
	ServerHTTPStart()
}
	

// ServerHTTPStart server http service
func ServerHTTPStart() (err error) {
	mux := http.NewServeMux()
	mux.HandleFunc(conf.ConstRouterGetObjectInfo, routerGetObjectInfo)
	mux.HandleFunc(conf.ConstRouterGetObjectConf, routerGetObjectConf)
	mux.HandleFunc(conf.ConstRouterGetObjectKeys, routerGetObjectKeys)
	mux.HandleFunc(conf.ConstRouterGetObjectDatas, routerGetObjectDatas)
	mux.HandleFunc(conf.ConstRouterGetObjectRestart, routerGetObjectRestart)
	mux.HandleFunc(conf.ConstRouterGetObjectAgain, routerGetObjectAgain)
	err = http.ListenAndServeTLS(":"+conf.Config[conf.ConstServerPort], "./cert/server.crt", "./cert/server.key", mux)
	return err
}
