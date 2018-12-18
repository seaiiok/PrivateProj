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

//3#门车辆磅秤
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

//检封离线数据压缩包
type clientsjfonlinefiles struct {
	clients body
}

func (p *clientsjfonlinefiles) ClientsProducterGetReady() error {
	err := mssql.GetDBInstance(p.clients.MyConf[Proto_SQL_clientsdbdriver], p.clients.MyConf[Proto_SQL_clientsdbconn])
	err = mssql.Ping()
	return err
}

func (p *clientsjfonlinefiles) ClientsProducterGetData() []string {
	res, err := mssql.DBExecSelectCol(p.clients.MyConf[Proto_SQL_clientsdbcol])
	if err != nil {
		return nil
	}
	return res
}

func (p *clientsjfonlinefiles) ClientsProducterGetDatas() [][]string {
	res, err := mssql.DBExecSelectRows(p.clients.MyConf[Proto_SQL_clientsdbrows])
	if err != nil {
		return nil
	}
	return res
}
