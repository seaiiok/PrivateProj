package net

import (
	"net/http"

	"ifix.cbpc/cbpc/pkg/conf"
)

// ServerHTTPStart server http service
func ServerHTTPStart() (err error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/routerSetObjectInfo", routerSetObjectInfo)
	mux.HandleFunc("/routerSetObjectConf", routerSetObjectConf)
	mux.HandleFunc("/routerSetObjectKeys", routerSetObjectKeys)
	mux.HandleFunc("/routerSetObjectDatas", routerSetObjectDatas)
	mux.HandleFunc("/routerSetObjectRestart", routerSetObjectRestart)
	err = http.ListenAndServeTLS(":"+conf.Config["serverport"], "./cert/server.crt", "./cert/server.key", mux)
	return err
}
