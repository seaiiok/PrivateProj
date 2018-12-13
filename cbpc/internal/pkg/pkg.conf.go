package pkg

import (
	"ifix.cbpc/cbpc/pkg/conf"
)

var config map[string]string

//预配置
func ConfigInit(path string) (err error) {
	if path=="" {
		path="./ifixtools.conf"
	}
	
	config=make(map[string]string,0)
	config,err=conf.GetConfKV(path)
	return
}