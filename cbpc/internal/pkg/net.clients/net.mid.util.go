package net

import (
	"errors"
	"fmt"

	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/print"
	"ifix.cbpc/cbpc/pkg/protocol"
)

func midController1thCom(p *protocol.Protocol) {

}

func midController2thCom(p *protocol.Protocol) {

}

func midController3thCom(p *protocol.Protocol) {

}

func midController4thCom(p *protocol.Protocol) string {
	sqlstr := ""
	for i := 0; i < len(p.Data); i++ {
		sqlstr = sqlstr + "'" + p.Data[i] + "',"
	}
	if sqlstr == "" {
		p.Error = errors.New("nothing data")
		return sqlstr
	}

	sqlstr = sqlstr[:len(sqlstr)-1]
	sqlstr = p.MeConf[conf.ConstClientsD3WeightRows] + " and F_ID in (" + sqlstr + ")"

	return sqlstr
}

func midController5thCom(p *protocol.Protocol) {

}

//GetMeTask ...
func GetMeTask(p *protocol.Protocol) {
	p.MeConf[conf.ConstClientsTask] = conf.Config[conf.ConstClientsTask]
	if p.MeConf[conf.ConstClientsTask] == "unkown" || p.MeConf[conf.ConstClientsTask] == "" {
		p.MeConf[conf.ConstClientsTask]= "unkown"
	}
}

type Usage struct {
	tip string
	key string
}

func (u *Usage) UserScanln() {
	u.tip = `请选择一个作业(数字)
1:3#门车辆称重系统  2:在线清数日志  3:检封离线压缩包  4:空调集中辅助`
	print.Line(print.Warn, u.tip)
	fmt.Scanln(&u.key)
	print.Line(print.Ok, "你选择了:"+u.key)
	if u.key != "1" && u.key != "2" && u.key != "3" && u.key != "4" {
		u.UserScanln()
	}
}
