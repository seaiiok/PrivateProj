package main

import (
	"fmt"

	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/mssql"
)

var config map[string]string

//预配置
func ConfigInit(path string) (err error) {
	if path == "" {
		path = "./ifixtools.conf"
	}

	config = make(map[string]string, 0)
	config, err = conf.GetConfKV(path)
	return
}
func main() {
	serverdbdriver := "odbc"
	serverdbconn := "driver={SQL Server};Server=192.168.1.18;uid=sa;pwd=kmtSoft12345678;port=1433;Database=CMTWeight121701"
	constrdemo := "select F_ID from T_Standard"
	db, _ := mssql.GetInstance(serverdbdriver, serverdbconn)
	res, _ := mssql.DBExecSelectCol(db, constrdemo)
	fmt.Println(res)

}
