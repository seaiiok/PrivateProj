package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"

	_ "github.com/alexbrainman/odbc" // mssql odbc driver
)

var msdb *sql.DB
var once sync.Once

//Init mssql driver init
func Init(drivername, datasourcename string) (err error) {
	// if me.db.Ping() == nil {
	// 	return nil
	// }
	once.Do(func() {
		msdb, err = sql.Open(drivername, datasourcename)
	})
	return msdb.Ping()
}

// Query  should query one col or panic
func Query(constr string) (res []string, err error) {
	res = make([]string, 0)
	stmt, err := msdb.Prepare(constr)
	fmt.Println("err1:", err)

	rows, err := stmt.Query()
	fmt.Println("err2:", err)
	fmt.Println("rows", rows)
	if err != nil || rows == nil {
		return
	}

	for rows.Next() {
		row := sql.NullString{}
		err = rows.Scan(&row)

		fmt.Println("err3:", err)

		if row.Valid {
			res = append(res, row.String)
		} else {
			res = append(res, "")
		}

	}

	defer func() {
		fmt.Println("defer1:", err)
		stmt.Close()
	}()

	defer func() {
		fmt.Println("defer2:", err)
		rows.Close()
	}()
	return
}

// Querys  query more
func Querys(constr string) (res [][]string, err error) {
	res = make([][]string, 0)
	stmt, err := msdb.Prepare(constr)
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
		re := make([]string, len(cols))
		for i := 0; i < len(cols); i++ {
			arr[i] = new(sql.NullString)
		}
		err = rows.Scan(arr...)
		if err != nil {
			return
		}

		for i := 0; i < len(arr); i++ {
			arrtemp := reflect.ValueOf(arr[i])
			arrtem := arrtemp.Interface().(*sql.NullString)
			re[i] = arrtem.String
		}
		res = append(res, re)
	}
	defer rows.Close()
	return
}

//Inserts insert more rows
func Inserts(constr string, res [][]string) (err error) {
	conn, err := msdb.Begin()

	for c := 0; c < len(res); c++ {
		re := make([]interface{}, len(res[c]))
		for i := 0; i < len(res[c]); i++ {
			re[i] = res[c][i]
		}
		stmt, _ := msdb.Prepare(constr)
		if err != nil {
			conn.Rollback()
		}
		defer stmt.Close()
		_, err := stmt.Exec(re...)
		if err != nil {
			conn.Rollback()
		}
		defer stmt.Close()
	}
	conn.Commit()
	return
}

func main() {
	// clientsdriver := "odbc"
	// clientssource := "driver={SQL Server};Server=192.168.1.18;uid=sa;pwd=kmtSoft12345678;port=1433;Database=CMTWeight121702"
	// err := Init(clientsdriver, clientssource)
	// fmt.Println(err)
	serverkeys := "a and b and c"
	// res, err := Query(serverkeys)
	// fmt.Println(err)
	// fmt.Println(res)
	// fmt.Println("exit ok")

	x2:=strings.TrimSpace(serverkeys)
	x2=strings.TrimSuffix(serverkeys, "c")
	fmt.Println(x2)
}

// serverkeys=select top 9000 F_ID from iFixsvr_D3Weight where F_IsFinish='1'  order by F_ID desc
