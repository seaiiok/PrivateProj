package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"sync"

	_ "github.com/alexbrainman/odbc" // mssql odbc driver
)

//DB sql struct
type DB struct {
	DriverName     string
	DataSourceName string
	sqlstr         string
	db             *sql.DB
	once           sync.Once
}

//Init mssql driver init
func (me *DB) Init() (err error) {
	// if me.db.Ping() == nil {
	// 	return nil
	// }
	me.once.Do(func() {
		me.db, err = sql.Open(me.DriverName, me.DataSourceName)
	})
	return me.db.Ping()
}

// Exec general
func (me *DB) Exec(constr string) error {
	stmt, err := me.db.Prepare(constr)
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
func (me *DB) Query(constr string) (res []string, err error) {
	res = make([]string, 0)
	stmt, err := me.db.Prepare(constr)
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
func (me *DB) Querys(constr string) (res [][]string, err error) {
	res = make([][]string, 0)
	stmt, err := me.db.Prepare(constr)
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
func (me *DB) Querys2Map(constr string) (res map[string]string, err error) {
	res = make(map[string]string, 0)
	stmt, err := me.db.Prepare(constr)
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
func (me *DB) Insert(constr string, res []string) (err error) {
	conn, err := me.db.Begin()
	re := make([]interface{}, len(res))
	for i := 0; i < len(res); i++ {
		re[i] = res[i]
	}
	stmt, err := me.db.Prepare(constr)
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

//Inserts insert more rows
func (me *DB) Inserts(constr string, res [][]string) (err error) {
	conn, err := me.db.Begin()

	for c := 0; c < len(res); c++ {
		re := make([]interface{}, len(res[c]))
		for i := 0; i < len(res[c]); i++ {
			re[i] = res[c][i]
		}
		stmt, err := me.db.Prepare(constr)
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
