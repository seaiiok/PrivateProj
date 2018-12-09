package conf

import (
"ifix.cbpc/cbpc/pkg/conf/config"
)

func ConfigGet(path string) (conf map[string]string, err error) {
	conf=make(map[string]string)
	conf,err=config.ConfGet(path)
	return
}

