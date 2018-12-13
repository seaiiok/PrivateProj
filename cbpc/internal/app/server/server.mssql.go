package server

import (
	"ifix.cbpc/cbpc/pkg/mssql"
	"fmt"
)

func DbInit(){
	db,_:=mssql.GetInstance(config["serverdbdriver"],config["serverdbconn"])
	res,_:=mssql.DBExecSelectCol(db,config["constr1"])
	fmt.Println(res)
}

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