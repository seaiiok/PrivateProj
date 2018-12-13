package pkg

import (
	"fmt"
	"net/http"
)

//测试路由
func serverTest(w http.ResponseWriter, r *http.Request) {
	var b []byte = []byte("测试路由")
	w.Write(b)
}

//内部路由
func serveriFixTools(w http.ResponseWriter, r *http.Request) {
	p := new(Proto)
	p.reader2struct(r.Body)
	p.header.ID = 100
	b := p.struct2arraybytes()
	w.Write(b)
	// p.servercontroller(w,r)
}

//管理路由
func serverOm(w http.ResponseWriter, r *http.Request) {
	//	serverPrint(w, r)
}

//打印输出  临时使用
func serverPrint(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Println(r.RemoteAddr)
	fmt.Println(r.Proto)
	fmt.Println(r.Method)
	w.Write([]byte("i am zognjian om"))
	fmt.Println(r.Header)
	return nil
}

func ServerHttpStart() {

	fmt.Println(config)

	mux := http.NewServeMux()
	mux.HandleFunc("/ifixtools", serveriFixTools)
	mux.HandleFunc("/om", serverOm)
	mux.HandleFunc("/test", serverTest)

	err := http.ListenAndServeTLS(":"+config["serverport"], "server.crt", "server.key", mux)
	if err != nil {
		fmt.Println(err)
	}
}
