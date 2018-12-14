package pkg

import (
	"net/http"
	"fmt"
)

//内部中间件
func (p *Proto)serveriFixToolsController(w http.ResponseWriter, r *http.Request) {
	p.reader2struct(r.Body)
	p.serveriFixToolsControllerTask(w)
}

//内部中间件
func (p *Proto)serveriFixToolsControllerTask(w http.ResponseWriter) {
	switch p.header.HeadMsg[Proto_Cmd] {

	case Proto_Cmd_Req_Ready:
	err:=ServerDBPing()
	if err != nil {
		p.header.HeadMsg[Proto_Cmd_Resp]=Proto_Cmd_Resp_No
		p.header.HeadMsg[Proto_Cmd]=Proto_Cmd_Resp_Again
		fmt.Println("---1---no")
	}else{
		p.header.HeadMsg[Proto_Cmd_Resp]=Proto_Cmd_Resp_Yes
		p.header.HeadMsg[Proto_Cmd]=Proto_Cmd_Req_KeyCols
		fmt.Println("---1---yes")
	}
	for k, v := range config {
		p.header.HeadMsg[k]=v
	}
	b := p.struct2arraybytes()
	w.Write(b)
	return

	case Proto_Cmd_Req_KeyCols:

		col,err:=dBExecSelectCol(p.header.HeadMsg[Proto_Conf_serverdbkeycol])
		if err != nil {
			p.header.HeadMsg[Proto_Cmd_Resp]=Proto_Cmd_Resp_No
			p.header.HeadMsg[Proto_Cmd]=Proto_Cmd_Resp_Again
			fmt.Println("---2---no")

		}else{
			p.header.HeadMsg[Proto_Cmd_Resp]=Proto_Cmd_Resp_Yes
			p.header.HeadMsg[Proto_Cmd]=Proto_Cmd_Req_InsertRows
			p.body.BodyData=col
			fmt.Println("---2---yes")
		}
		b := p.struct2arraybytes()
		w.Write(b)
		return

	case Proto_Cmd_Req_InsertRows:
		res,_:=dBExecSelectRows(p.header.HeadMsg[Proto_Conf_clientsd3weightselectrows])
		fmt.Println(res)
		err:=dBExecInsertRow(p.header.HeadMsg[Proto_Conf_clientsd3weightinsertrows],res[0])
		if err != nil {
			p.header.HeadMsg[Proto_Cmd_Resp]=Proto_Cmd_Resp_No
			p.header.HeadMsg[Proto_Cmd]=Proto_Cmd_Resp_Again
			fmt.Println("---3---no")

		}else{
			p.header.HeadMsg[Proto_Cmd_Resp]=Proto_Cmd_Resp_Yes
			p.header.HeadMsg[Proto_Cmd]=Proto_Cmd_Req_InsertRows
			fmt.Println("---3---yes")
		}
		b := p.struct2arraybytes()
		w.Write(b)
		return

	case Proto_Cmd_Resp_Again:
		err:=ServerDBPing()
		if err != nil {
			p.header.HeadMsg[Proto_Cmd_Resp]=Proto_Cmd_Resp_No
			p.header.HeadMsg[Proto_Cmd]=Proto_Cmd_Resp_Again
			fmt.Println("---a---no")
		}else{
			p.header.HeadMsg[Proto_Cmd_Resp]=Proto_Cmd_Resp_Yes
			p.header.HeadMsg[Proto_Cmd]=Proto_Cmd_Req_KeyCols
			fmt.Println("---a---yes")
		}
		for k, v := range config {
			p.header.HeadMsg[k]=v
		}
		b := p.struct2arraybytes()
		w.Write(b)
		return
	default:
				
	}
	

}
