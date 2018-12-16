package pkg

import (
	"ifix.cbpc/cbpc/pkg/log"
	"net/http"
)

//测试路由
func serverTest(w http.ResponseWriter, r *http.Request) {
}

//内部路由
func serveriFixTools(w http.ResponseWriter, r *http.Request) {
	p := NewProto()
    p.serveriFixToolsController(w,r)
}

//管理路由
func serverOm(w http.ResponseWriter, r *http.Request) {
	//	serverPrint(w, r)
}

// ServerHTTPStart http service
func ServerHTTPStart() {

	mux := http.NewServeMux()
	mux.HandleFunc("/ifixtools", serveriFixTools)
	mux.HandleFunc("/om", serverOm)
	mux.HandleFunc("/test", serverTest)

	err := http.ListenAndServeTLS(":"+config["serverport"], "./cert/server.crt", "./cert/server.key", mux)
	if err != nil {
		log.Log(err)
	}
}
