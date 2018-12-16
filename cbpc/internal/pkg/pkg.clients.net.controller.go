package pkg

import (
	"fmt"
	"time"

	"ifix.cbpc/cbpc/internal/pkg/pkg.producter"
)

var icp clientsproducter.IProducte
var clientsconfig map[string]string

func clientsproducterparams(col, rows string) {
	clientsd3weightparams := new(clientsproducter.ClientsD3weight)
	clientsd3weightparams.Clientsd3weightdbdriver = clientsconfig["clientsd3weightdbdriver"]
	clientsd3weightparams.Clientsd3weightdbconn = clientsconfig["clientsd3weightdbconn"]
	clientsd3weightparams.Clientsd3weightkeycol = col
	clientsd3weightparams.Clientsd3weightselectrows = rows
	icp = clientsd3weightparams
}

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

		p.header.HeadMsg[Proto_Conf_clientstask] = clientsconfig["clientstask"]

		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_1th
		//准备
	case Proto_Cmd_2th:
		clientsconfig = make(map[string]string, 0)
		clientsconfig = p.header.HeadMsg
		clientsproducterparams("", "")
		err := icp.ClientsProducterGetReady()
		if err != nil {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Restart
		}

		//获取新增数据
	case Proto_Cmd_3th:
		col := SqlStringMakeDTime(clientsconfig["clientsd3weightkeycol"])
		clientsproducterparams(col, "")
		res := icp.ClientsProducterGetData()
		p.body.BodyData = make([]string, 0)
		if len(res) > 0 {
			p.body.BodyData = res
		} else {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Again
		}

	case Proto_Cmd_4th:
		rows := sqlStringMakeRows(clientsconfig["clientsd3weightselectrows"], p.body.BodyData)
		clientsproducterparams("", rows)
		res := icp.ClientsProducterGetDatas()
		p.body.BodyDatas = make([][]string, 0)
		if len(res) > 0 {
			for i := 0; i < len(res); i++ {
				p.body.BodyDatas = append(p.body.BodyDatas, res[i])
			}
		}

	case Proto_Cmd_5th:
		time.Sleep(1 * time.Second)
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_2th
	case Proto_Cmd_Restart:
		time.Sleep(1 * time.Second)
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_2th
	case Proto_Cmd_Again:
		time.Sleep(1 * time.Second)
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_2th
	default:

	}

}
