package service

import (
	"fmt"
	"net/http"
)

//测试路由
func serverTest(w http.ResponseWriter, r *http.Request){
	p:=new(MyProto)
	p.reader2struct(r.Body)
	p.servercontroller(w,r)
}

//内部路由
func serveriFixTools(w http.ResponseWriter, r *http.Request){
    // serverPrint(w,r)
	p:=new(MyProto)
	p.reader2struct(r.Body)
	fmt.Println(p)
	p.Header.Id=100
	b:=p.struct2arraybytes()
	w.Write(b)
	// p.servercontroller(w,r)
}

//管理路由
func serverOm(w http.ResponseWriter, r *http.Request){
	serverPrint(w,r)
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

func ServerOpen() {

	conf,err:=ConfigMap()
	fmt.Println(conf)

	


	mux:=http.NewServeMux()
	mux.HandleFunc("/ifixtools",serveriFixTools)
	mux.HandleFunc("/om",serverOm)
	mux.HandleFunc("/test",serverTest)
	
	err=http.ListenAndServeTLS(":"+conf["ServerPort"], "server.crt", "server.key", mux)
	if err != nil {
	fmt.Println(err)
	}
}


