package pkg

import (
	"fmt"

	"ifix.cbpc/cbpc/pkg/mssql"

	//	"ifix.cbpc/cbpc/pkg/mssql"
)

//type Ms struct {
//	D3Weight
//}
//	ServerDatebaseGetReady(req *Proto) *Proto
//	ServerDatebaseGetData(req *Proto) *Proto
//	ServerDatebaseSetData(req *Proto) *Proto

func (m *D3Weight) ServerDatebaseGetReady(req *Proto) *Proto {
	db, _ := mssql.GetInstance(config["serverdbdriver"], config["serverdbconn"])
	res, _ := mssql.DBExecSelectCol(db, config["constrdemo"])
	fmt.Println(res)
	return req
}

func (m *D3Weight) ServerDatebaseGetData(req *Proto) *Proto {
	return req
}

func (m *D3Weight) ServerDatebaseSetData(req *Proto) *Proto {
	return req
}

//func (s *IServerDatebase)dbready
// (config["serverdbdriver"],config["serverdbconn"])
// driver:="odbc"
// conn:="driver={SQL Server};Server=192.168.1.18;uid=sa;pwd=kmtSoft12345678;port=1433;Database=CMTWeight121701"
// msdb,err:=db.GetInstance(driver,conn)
// if err != nil {
// 	fmt.Println(err)
// }
// res,err:=db.DBExecSelectCols(msdb,"select F_ID from T_Standard")
// if err != nil {
// 	fmt.Println(err)
// }
// fmt.Println(res)