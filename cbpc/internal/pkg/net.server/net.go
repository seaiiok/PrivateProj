package net

import (
	"net/http"

	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/convert"
	"ifix.cbpc/cbpc/pkg/protocol"
)

//serverD3Weight router
func serverRouter(w http.ResponseWriter, r *http.Request) {
	p := protocol.NewProto()
	err := convert.Reader2Struct(r.Body, p)
	if err == nil {
		MidController(w, p)
	}
}

// ServerHTTPStart http service
func ServerHTTPStart() (err error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/service", serverRouter)
	err = http.ListenAndServeTLS(":"+conf.Config["serverport"], "./cert/server.crt", "./cert/server.key", mux)
	return err
}
