package main

import (
	"fmt"
	"database/sql"
	_ "github.com/alexbrainman/odbc"

)
var db *sql.DB

//sql server 驱动驱动初始化
func MssqlInit(driverName, dataSourceName string) (err error) {
	db,err=DBInit(driverName, dataSourceName)
	return err
}

//数据库驱动初始化
func DBInit(driverName, dataSourceName string) (db *sql.DB,err error) {
	db, err = sql.Open(driverName, dataSourceName)
	err = db.Ping()
	return
}

// 一般事务
func DBExec(db *sql.DB,constr string) error {
	stmt, err := db.Prepare(constr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func main(){
	MssqlInit("odbc",)
}



