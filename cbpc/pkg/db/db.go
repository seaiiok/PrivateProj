package db

import (
	"database/sql"
	"reflect"
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

// Exec general
func Exec(constr string) error {
	stmt, err := msdb.Prepare(constr)
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

// Query  should query one col or panic
func Query(constr string) (res []string, err error) {
	res = make([]string, 0)
	stmt, err := msdb.Prepare(constr)
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
		} else {
			res = append(res, "")
		}

	}
	defer rows.Close()
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

//Querys2Map only query like key value
func Querys2Map(constr string) (res map[string]string, err error) {
	res = make(map[string]string, 0)
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
		// re := make([]string,len(cols))
		for i := 0; i < len(cols); i++ {
			arr[i] = new(sql.NullString)
		}
		err = rows.Scan(arr...)
		if err != nil {
			return
		}
		if len(arr) == 2 {
			ktemp := reflect.ValueOf(arr[0])
			ktem := ktemp.Interface().(*sql.NullString)
			vtemp := reflect.ValueOf(arr[1])
			vtem := vtemp.Interface().(*sql.NullString)
			res[ktem.String] = vtem.String
		}
	}
	defer rows.Close()
	return
}

//Insert insert one row
func Insert(constr string, res []string) (err error) {
	conn, err := msdb.Begin()
	re := make([]interface{}, len(res))
	for i := 0; i < len(res); i++ {
		re[i] = res[i]
	}
	stmt, err := msdb.Prepare(constr)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(re...)
	if err != nil {
		conn.Rollback()
		return
	}
	conn.Commit()
	defer stmt.Close()
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
