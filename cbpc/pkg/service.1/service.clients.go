package service

import (
	"fmt"
)

func ClientsOpen() {
	p:=new(MyProto)
	p=ClientsProtoInit()
	
	p.clientInit()
	p.Header.Identify=Proto_Cmd_Ok
	p.Header.Id=Proto_Id_1st
	p.httpPost()
	fmt.Println(p)
}

