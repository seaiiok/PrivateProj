package service

import (
	"fmt"
	"net/http"
)

func (p *MyProto) servercontroller(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
	p.reader2struct(r.Body)
	if p.flowstep() {
		p.ServerProtoInit()	
		w.Write(p.struct2arraybytes())	
		return 
	}
	switch p.Header.Id {
	case Proto_Id_init:
		//客服端请求初始化
		p.ServerProtoInit()	
	case Proto_Id_1st:
		//客服端请求分配任务
		
		// p.Header.Id=p.Header.Id+1
		p.Header.HeadMsg[Proto_Map_ClientTask]="my job is d3weight"
	case Proto_Id_2nd:
		//服务器数据库OK？
		p.Header.HeadMsg[Proto_Map_ClientTask]="my job is offline"
	case Proto_Id_3rd:
		//服务器数据库OK？
	}
	w.Write(p.struct2arraybytes())
}

//流程机
//未通过，则要求重新开始流程
func (p *MyProto) flowstep() bool {
	if p.Header.Identify != Proto_Cmd_Ok {
		return true
	}
	if p.Header.Method != "POST" {
		return true
	}
	return false
}
//客服端请求
func (p *MyProto) serverResponseClientTask() (MyProto,error) {

	p.Header.HeadMsg[Proto_Map_ClientTask]=""
	
	return *p,nil
}



