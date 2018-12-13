package pkg

import (
	"net/http"
)

const (
	Proto_Errors = iota
	Proto_Id_init
	Proto_Id_Rdy
	Proto_Id_Nil
	Proto_Id_Err
	Proto_Id_Start
	Proto_Id_Stop
	Proto_Id_Restart
	Proto_Id_1st
	Proto_Id_2nd
	Proto_Id_3rd
	Proto_Id_4th
	Proto_Id_5th
	Proto_Id_6th
	Proto_Id_7th
	Proto_Id_8th
	Proto_Id_9th
	Proto_Id_10th
	Proto_Id_11th
	Proto_Id_12th
)

const (
	Proto_Cmd_Ok                   = true
	Proto_Cmd_No                   = false
	Proto_Map_ServerDriver         = "金刚芭比-小哪吒"
	Proto_Map_ServerConn           = "是他,是他,是他,就是他,我们的朋友小哪吒"
	Proto_Map_ServerTableName      = "是他,就是他,是他,就是他,少年英雄小哪吒"
	Proto_Map_ServerTableCol       = "上天他比天要高,下海他比海更大啊~啊~"
	Proto_Map_ServerSqlCreateTable = "智斗妖魔勇降鬼怪,少年英雄就是小哪吒"
	Proto_Map_ServerSqlCol         = "有时他很聪明,有时他也犯傻,他的个头跟我一般高"
	Proto_Map_ServerSqlExistTable  = "有时他很努力,有时他也贪玩,他的年纪和我一般大"
	Proto_Map_ClientDriver         = "啦啦啦啦~啦啦啦啦~1"
	Proto_Map_ClientConn           = "啦啦啦啦~啦啦啦啦~2"
	Proto_Map_ClientTableName      = "啦啦啦啦~啦啦啦啦~3"
	Proto_Map_ClientTableCol       = "啦啦啦啦~啦啦啦啦~4"
	Proto_Map_ClientSqlCol         = "啦啦啦啦~啦啦啦啦~5"
	Proto_Map_ClientSqlExistTable  = "啦啦啦啦~啦啦啦啦~6"
	Proto_Map_ClientSql            = "啦啦啦啦~啦啦啦啦~7"
	Proto_Map_ClientTask           = "啦啦啦啦~啦啦啦啦~8"
)

var client *http.Client

// var p *MyProto=new(MyProto)

type header struct {
	//流水号
	ID int
	//请求Url
	URL string
	//请求方式
	Method string
	//本机ip地址
	Ipadrr string
	//验证
	Identify bool
	//数据条数
	Count int
	//数据结构
	Dstruct string
	//错误
	Error error
	//流程追溯
	DataTrace []string
	//请求资源
	HeadMsg map[string]string
}

type body struct {
	BodyMsg map[int]map[string]string
}

type Proto struct {
	header
	body
}
