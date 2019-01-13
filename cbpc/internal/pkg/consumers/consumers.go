package consumers

import (
	"ifix.cbpc/cbpc/pkg/db"
	"ifix.cbpc/cbpc/pkg/protocol"
	"ifix.cbpc/cbpc/pkg/utils"
)

//Objects ...
type Objects struct {
	protocol.Proto
}

//GetObjectInfo ...
func (p *Objects) GetObjectInfo() {

}

//GetObjectConf ...
func (p *Objects) GetObjectConf() {
	err := db.Init(p.SQL.DatabaseDriver, p.SQL.DatabaseSource)
	if err != nil {
		p.SetError(true)
	}
}

//GetObjectKeys ...
func (p *Objects) GetObjectKeys() {
	res := make([]string, 0)
	res, err := db.Query(p.SQL.QuerySQL)
	if err != nil {
		p.SetError(true)
	}
	r := utils.ArrayDiff(p.SQL.Args, res)
	p.SetArgs(r)
}

//GetObjectDatas ...
func (p *Objects) GetObjectDatas() {
	err := db.Inserts(p.SQL.InsertSQL, p.SQL.Data)
	if err != nil {
		p.SetError(true)
	}
}
