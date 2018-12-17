package pkg

import (
	"fmt"
	"net/http"
)

//内部中间件
func (p *Proto) serveriFixToolsController(w http.ResponseWriter, r *http.Request) {
	p.reader2struct(r.Body)
	p.serveriFixToolsControllerTask(w)
}

//内部中间件
func (p *Proto) serveriFixToolsControllerTask(w http.ResponseWriter) {
	p.ProcessTrace(p.header.HeadMsg[Proto_Cmd])
	fmt.Println(p.header.ProcessTrace)
	switch p.header.HeadMsg[Proto_Cmd] {
	//客服端的任务确认
	case Proto_Cmd_1th:

		err := p.GetServerConf()
		if err != nil {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Restart
		} else {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_2th
		}
		b := p.struct2arraybytes()
		w.Write(b)
		return

		//服务端数据库自检
	case Proto_Cmd_2th:

		err := p.ServerDBPing()
		if err != nil {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Restart
		} else {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_3th
		}

		b := p.struct2arraybytes()
		w.Write(b)
		return

	//开始任务
	case Proto_Cmd_3th:
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_4th
		b := p.struct2arraybytes()
		w.Write(b)
		return

	case Proto_Cmd_4th:

		col, err := p.ServerDBExecSelectCol()
		fmt.Println(col)
		resarr := ArrayDiff(p.body.BodyData, col)

		if err != nil || len(resarr) == 0 {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Again

		} else {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_5th
			p.body.BodyData = make([]string, 0)
			p.body.BodyData = resarr
		}

		b := p.struct2arraybytes()
		w.Write(b)
		return

	case Proto_Cmd_5th:

		err := p.ServerDBExecInsertRows()
		if err != nil {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Again
		} else {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_6th
		}

		b := p.struct2arraybytes()
		w.Write(b)
		return

	case Proto_Cmd_Restart:
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_1th
		b := p.struct2arraybytes()
		w.Write(b)
		return

	case Proto_Cmd_Again:
		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_2th
		b := p.struct2arraybytes()
		w.Write(b)
		return

	default:

		b := p.struct2arraybytes()
		w.Write(b)
		return
	}

}
