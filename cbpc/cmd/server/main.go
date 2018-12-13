package main

import (
	"time"
	"ifix.cbpc/cbpc/internal/app/server"
	"ifix.cbpc/cbpc/pkg/print"
)

func main(){
	err:=server.ConfigInit("")
	if err != nil {
		print.Line(print.Err,"预读配置失败,工具退出")
		return
	}
	print.Line(print.Ok,"预读配置成功")
	server.DbInit()
	time.Sleep(2*time.Second)
}