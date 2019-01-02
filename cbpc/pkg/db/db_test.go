package db

import "testing"

func TestInit(t *testing.T) {
d:=new(DB)
d.DriverName="odbc"
d.DataSourceName="driver={SQL Server};Server=192.168.1.18;uid=sa;pwd=kmtSoft12345678;port=1433;Database=CMTWeight121701"
d.Init()
}
