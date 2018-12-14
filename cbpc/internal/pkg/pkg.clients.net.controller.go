package pkg

import (
	"fmt"
	"time"
)

//内部中间件
func (p *Proto) clientsiFixToolsController() {
	//发送job
	p.header.HeadMsg[Proto_Cmd]=Proto_Cmd_Req_Ready
	p.httpPost()
	for {
		time.Sleep(15 * time.Second)
		p.httpPost()
		p.clientsiFixToolsControllerTask()
	}

}

//内部中间件
func (p *Proto) clientsiFixToolsControllerTask() {
	switch p.header.HeadMsg[Proto_Cmd] {

	case Proto_Cmd_Req_Ready:
		fmt.Println("------1---------")
	   fmt.Println(p.header.HeadMsg[Proto_Cmd_Resp])
	   fmt.Println(p.header.HeadMsg[Proto_Cmd])

	case Proto_Cmd_Req_KeyCols:

		fmt.Println("------2---------")
	   fmt.Println(p.header.HeadMsg[Proto_Cmd_Resp])
	   fmt.Println(p.header.HeadMsg[Proto_Cmd])

	case Proto_Cmd_Req_InsertRows:

		fmt.Println("------3---------")
	   fmt.Println(p.header.HeadMsg[Proto_Cmd_Resp])
	   fmt.Println(p.header.HeadMsg[Proto_Cmd])

	case Proto_Cmd_Resp_Again:
		fmt.Println("------4---------")
		fmt.Println(p.header.HeadMsg[Proto_Cmd_Resp])
		fmt.Println(p.header.HeadMsg[Proto_Cmd])
	default:
		fmt.Println("------5---------")
		fmt.Println(p.header.HeadMsg[Proto_Cmd_Resp])
		fmt.Println(p.header.HeadMsg[Proto_Cmd])
	}

}
