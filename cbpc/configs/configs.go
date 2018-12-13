package configs

import (
	"ifix.cbpc/cbpc/pkg/conf"
)

func ConfigInit(path string) (config map[string]string, err error) {
	if path == "" {
		path = "./ifixtools.conf"
	}
	config = make(map[string]string, 0)
	config, err = conf.GetConfKV(path)
	return
}
