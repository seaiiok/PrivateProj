package clientsproducter

import (
	"ifix.cbpc/cbpc/pkg/mssql"
)

type IProducte interface {
	ClientsProducterGetReady() error
	ClientsProducterGetData() []string
	ClientsProducterGetDatas() [][]string
}

type ClientsD3weight struct {
	Clientsd3weightdbdriver   string
	Clientsd3weightdbconn     string
	Clientsd3weightkeycol     string
	Clientsd3weightselectrows string
}

func (p *ClientsD3weight) ClientsProducterGetReady() error {
	err := mssql.GetDBInstance(p.Clientsd3weightdbdriver, p.Clientsd3weightdbconn)
	return err
	err = mssql.Ping()
	return err
}

func (p *ClientsD3weight) ClientsProducterGetData() []string {
	res, err := mssql.DBExecSelectCol(p.Clientsd3weightkeycol)
	if err != nil {
		return nil
	}
	return res
}

func (p *ClientsD3weight) ClientsProducterGetDatas() [][]string {
	res, err := mssql.DBExecSelectRows(p.Clientsd3weightselectrows)
	if err != nil {
		return nil
	}
	return res
}
