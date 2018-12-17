package pkg

import (
	"ifix.cbpc/cbpc/pkg/mssql"
)

// serverdbdriver serverdbconn serverdbcol insertdbrows clientsdbdriver clientsdbconn clientsdbcol clientsdbrows

func ServerDBInit() error {
	return mssql.GetDBInstance(config[Proto_SQL_serverdbdriver], config[Proto_SQL_serverdbconn])
}

func (c *Proto) ServerDBPing() error {
	return mssql.Ping()
}

func (c *Proto) ServerDBExecSelectCol() (res []string, err error) {
	c.sqlStringMakeDTimes(config[Proto_SQL_serverdbcol])
	return mssql.DBExecSelectCol(c.MyConf[Proto_SQL_serverdbcol])
}

func (c *Proto) ServerDBExecInsertRows() (err error) {
	c.sqlStringServerMakeRows(config[Proto_SQL_serverdbrows])
	return mssql.DBExecInsertRows(c.MyConf[Proto_SQL_serverdbrows], c.BodyDatas)
}

func (c *Proto) ClientsDBInit() error {
	return mssql.GetDBInstance(c.MyConf[Proto_SQL_clientsdbdriver], c.MyConf[Proto_SQL_clientsdbconn])
}

func (c *Proto) ClientsDBPing() error {
	return mssql.Ping()
}

func (c *Proto) ClientsDBExecSelectCol() (res []string, err error) {
	return mssql.DBExecSelectCol(c.body.MyConf[Proto_SQL_clientsdbcol])
}

func (c *Proto) ClientsDBExecSelectRows() (res [][]string, err error) {
	return mssql.DBExecSelectRows(c.body.MyConf[Proto_SQL_clientsdbrows])
}

// func (c *body)DBExecInsertRow(constr string, res []string) (err error){
// 	return mssql.DBExecInsertRow(constr, res)
// }
