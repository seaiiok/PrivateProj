package mssql

import (
	"fmt"
	"sync"
	"database/sql"
	_ "github.com/alexbrainman/odbc"
	"reflect"
)
var db *sql.DB
var once sync.Once


//数据库驱动初始化--mssql
func GetDBInstance(driverName, dataSourceName string) error {
	var err error
    once.Do(func() {
		db,err=sql.Open(driverName, dataSourceName)
    })
	return err
}


func Ping() error {
return db.Ping()
}

// 一般事务
func DBExec(constr string) error {
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
func DBExecSelectCol(constr string) (res []string, err error) {
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
func DBExecSelectRows(constr string) (res [][]string, err error) {
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
// 多条记录查询
func DBExecSelectRows1(constr string) (res [][]string, err error) {
	fmt.Println(constr)
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
fmt.Println(arr)
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
func DBExecSelectMap4Config(constr string) (res map[string]string, err error) {
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
func DBExecInsertRow(constr string,res []string) (err error) {
	conn, err := db.Begin()
	re:=make([]interface{},len(res))
	for i := 0; i < len(res); i++ {
		re[i] = res[i]
	}
	stmt, err := db.Prepare(constr)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Exec(re...)
	if err != nil {
		conn.Rollback()
		return
	}
	fmt.Println(rows)
	conn.Commit()
	defer stmt.Close()
	return
}

// 插入所有列数据
// 一般事务
func DBExecInsertRows(constr string,res [][]string) (err error) {
	conn, err := db.Begin()

	for c := 0; c < len(res); c++ {
		re:=make([]interface{},len(res[c]))
		for i := 0; i < len(res[c]); i++ {
			re[i] = res[c][i]
		}
		stmt, err := db.Prepare(constr)
		if err != nil {
			fmt.Println(err)
		}
		defer stmt.Close()
		rows, err := stmt.Exec(re...)
		if err != nil {
			conn.Rollback()
			fmt.Println(err)
		}
		fmt.Println(rows)
		defer stmt.Close()
	}
	conn.Commit()
	return
}
