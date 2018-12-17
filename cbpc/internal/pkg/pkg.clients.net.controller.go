package pkg

import (
	"fmt"
	"time"
)

//内部中间件
func (p *Proto) clientsiFixToolsController() {
	p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_1th

	for {
		p.clientsiFixToolsControllerTask()
		p.httpPost()
	}

}

//内部中间件
func (p *Proto) clientsiFixToolsControllerTask() {
	fmt.Println(p.header.HeadMsg[Proto_Cmd])
	switch p.header.HeadMsg[Proto_Cmd] {

	//声明职责
	case Proto_Cmd_1th:

		p.header.HeadMsg[Proto_Conf_clientstask] = config[Proto_Conf_clientstask]
		//准备
	case Proto_Cmd_2th:
		p.GetClientsConf()

	case Proto_Cmd_3th:

		//获取新增数据
	case Proto_Cmd_4th:
		
		err := p.GetClientsCol()
		if err != nil {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Again
		}

	case Proto_Cmd_5th:
		fmt.Println()
		err := p.GetClientsRows()
		if err != nil {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Again
		}

	case Proto_Cmd_6th:
		time.Sleep(1 * time.Second)
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_3th
		fmt.Println("6")
	case Proto_Cmd_Restart:
		time.Sleep(1 * time.Second)
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_3th
		fmt.Println("r")
	case Proto_Cmd_Again:
		time.Sleep(1 * time.Second)
		fmt.Println("a")
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_3th
	default:

	}

}
