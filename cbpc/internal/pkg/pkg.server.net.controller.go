package pkg

import (
	"fmt"
	"net/http"
)
type tempsqlstr struct {
	tempsqlstr1 string
	tempsqlstr2 string
	tempsqlstr3 string
}
//内部中间件
func (p *Proto) serveriFixToolsController(w http.ResponseWriter, r *http.Request) {
	p.reader2struct(r.Body)
	p.serveriFixToolsControllerTask(w)
}

//内部中间件
func (p *Proto) serveriFixToolsControllerTask(w http.ResponseWriter) {
	fmt.Println(p.header.ProcessTrace)
	p.ProcessTrace(p.header.HeadMsg[Proto_Cmd])
	switch p.header.HeadMsg[Proto_Cmd] {
	//客服端的任务确认
	case Proto_Cmd_1th:
		if p.header.HeadMsg[Proto_Conf_clientstask] != "unknown" {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_2th
		} else {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Restart
		}

		for k, v := range config {
			p.header.HeadMsg[k] = v
		}

		b := p.struct2arraybytes()
		w.Write(b)
		return

		//服务端数据库自检
	case Proto_Cmd_2th:

		err := ServerDBPing()
		if err != nil {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Restart
		} else {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_3th
		}

		b := p.struct2arraybytes()
		w.Write(b)
		return

	case Proto_Cmd_3th:
		col, err := dBExecSelectCol(SqlStringMakeDTime(config[Proto_Conf_serverd3weightkeycol]))
		fmt.Println(p.body.BodyData)
		fmt.Println(col)
		resarr := ArrayDiff(p.body.BodyData,col)
		fmt.Println(resarr)
		if err != nil || len(resarr)==0 {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Again

		} else {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_4th
			p.body.BodyData = make([]string, 0)
			p.body.BodyData = resarr
		}

		b := p.struct2arraybytes()
		w.Write(b)
		return

	case Proto_Cmd_4th:

		err := dBExecInsertRows(config[Proto_Conf_serverd3weightinsertrows], p.body.BodyDatas)
		if err != nil {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Again
		} else {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_5th
		}

		b := p.struct2arraybytes()
		w.Write(b)
		return

	default:

	}

}

func (p *Proto)dBExecSelectColMapClient() string {
	tss:=new(tempsqlstr)
	switch p.header.HeadMsg[Proto_Conf_clientstask] {
	case Proto_Conf_Task_d3weight:
		tss.tempsqlstr1=""
	}
	if p.header.HeadMsg[Proto_Conf_clientstask] != "unknown" {
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_2th
	} else {
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Restart
	}
	return tss
}
