package service

import (
	"fmt"
	"database/sql"
	_ "ii.sea/libs/odbc"
	_ "ii.sea/libs/SQLite"
	"reflect"
)
var db *sql.DB
var db3 *sql.DB

//sql server 驱动驱动初始化
func MssqlInit(driverName, dataSourceName string) (err error) {
	db,err=DBInit(driverName, dataSourceName)
	return err
}

//sqlite数据库驱动初始化
func SQLiteInit(driverName, dataSourceName string) (err error) {
	db3,err=DBInit(driverName, dataSourceName)
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

// 查询-Keys
// 单列记录查询 否则panic
func DBExecSelectCol(db *sql.DB,constr string) (res []string, err error) {
	res=make([]string,0)
	stmt, err := db.Prepare(constr)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, _ := stmt.Query()
	for rows.Next() {
		row := sql.NullString{}
		err = rows.Scan(&row)
		if err != nil {
			return
		}
		if row.Valid {
			res = append(res, row.String)
		}else{
			res = append(res, "")
		}

	}
	defer rows.Close()
	return
}

// 查询-*
// 多条记录查询
func DBExecSelectCols(db *sql.DB,constr string) (res [][]string, err error) {
	res=make([][]string,0)
	stmt, err := db.Prepare(constr)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, _ := stmt.Query()
	cols, err := rows.Columns()
	if err != nil {
		return
	}
	for rows.Next() {
	arr := make([]interface{}, len(cols))
	re := make([]string,len(cols))
	for i := 0; i < len(cols); i++ {
		arr[i] = new(sql.NullString)
	}
		err = rows.Scan(arr...)
		if err != nil {
			return
		}

	for i := 0; i < len(arr); i++ {
		_arr := reflect.ValueOf(arr[i])
		__arr:=_arr.Interface().(*sql.NullString)
		re[i]=__arr.String
	}
		res=append(res,re)	
	}
	defer rows.Close()
	return
}

// 查询-*
// 配置表
func DBExecSelectMap4Config(db *sql.DB,constr string) (res map[string]string, err error) {
	res=make(map[string]string,0)
	stmt, err := db.Prepare(constr)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, _ := stmt.Query()
	cols, err := rows.Columns()
	if err != nil {
		return
	}
	for rows.Next() {
	arr := make([]interface{}, len(cols))
	// re := make([]string,len(cols))
	for i := 0; i < len(cols); i++ {
		arr[i] = new(sql.NullString)
	}
		err = rows.Scan(arr...)
		if err != nil {
			return
		}
	if len(arr)==2 {
		_k := reflect.ValueOf(arr[0])
		__k:=_k.Interface().(*sql.NullString)
		_v := reflect.ValueOf(arr[1])
		__v:=_v.Interface().(*sql.NullString)
		res[__k.String]=__v.String
	}
	}
	defer rows.Close()
	return
}

// 插入所有列数据
// 一般事务
func DBExecInsertCols(db *sql.DB,constr string,res []string) (err error) {
	conn, err := db.Begin()
	r1:=make([]interface{},len(res))
	for i := 0; i < len(res); i++ {
		r1[i] = res[i]
	}
	stmt, err := db.Prepare(constr)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Exec(r1...)
	if err != nil {
		conn.Rollback()
		return
	}
	fmt.Println(rows)
	conn.Commit()
	defer stmt.Close()
	return
}
