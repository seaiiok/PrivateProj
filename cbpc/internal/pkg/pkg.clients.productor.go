package pkg

import (
	"ifix.cbpc/cbpc/pkg/mssql"
)

//clients data interface
type IProducte interface {
	ClientsProducterGetReady() error
	ClientsProducterGetData() []string
	ClientsProducterGetDatas() [][]string
}

type clientsd3weight struct {
	clients body
}

func (p *clientsd3weight) ClientsProducterGetReady() error {
	err := mssql.GetDBInstance(p.clients.MyConf[Proto_SQL_clientsdbdriver], p.clients.MyConf[Proto_SQL_clientsdbconn])
	err = mssql.Ping()
	return err
}

func (p *clientsd3weight) ClientsProducterGetData() []string {
	res, err := mssql.DBExecSelectCol(p.clients.MyConf[Proto_SQL_clientsdbcol])
	if err != nil {
		return nil
	}
	return res
}

func (p *clientsd3weight) ClientsProducterGetDatas() [][]string {
	res, err := mssql.DBExecSelectRows(p.clients.MyConf[Proto_SQL_clientsdbrows])
	if err != nil {
		return nil
	}
	return res
}
