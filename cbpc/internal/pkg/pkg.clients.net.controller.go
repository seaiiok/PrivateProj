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
		err:=p.reconnect()
		if err != nil {
			return
		}
	}

}

//内部中间件
func (p *Proto) clientsiFixToolsControllerTask() {
	p.ProcessTrace(p.header.HeadMsg[Proto_Cmd])
	fmt.Println(p.header.ProcessTrace)
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
		err := p.GetClientsRows()
		if err != nil {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Again
		}

	case Proto_Cmd_6th:
		time.Sleep(Proto_Conf_ProcessSleep * time.Second)
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_3th
	case Proto_Cmd_Restart:
		time.Sleep(Proto_Conf_ProcessSleep * time.Second)
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_1th
	case Proto_Cmd_Again:
		time.Sleep(Proto_Conf_ProcessSleep * time.Second)
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_3th
	default:

	}

}

func (p *Proto) reconnect() error {
	if p.header.Error != nil {
		time.Sleep(Proto_Conf_ProcessSleep * time.Second)
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Restart
		ClientsHTTPStart()
	}

	return p.header.Error
}
