package pkg

import (
	"fmt"
	"ifix.cbpc/cbpc/pkg/conf"
)

var config map[string]string

//服务端预配置
func ServerConfigInit(path string) (err error) {
	if path == "" {
		path = "ifixtools.conf"
	}

	config = make(map[string]string, 0)
	config, err = conf.GetConfKV(path)
	return
}

//客服端预配置
func ClientsConfigInit(path string) (err error) {
	if path == "" {
		path = "ifixtools.conf"
	}
	err = ServerConfigInit(path)
	if config["clientstask"] == "" || config["clientstask"] == "unkown" {
		config["clientstask"] = ClientStartInput()
		err = ConfigSet(path, config)
	}
	return
}

//保持
func ConfigSet(path string, config map[string]string) (err error) {
	fmt.Println(config)
	err=conf.SetConfKV(path, config)
	return 
}

// if pkg.clientsconfig["clientstask"] !="" || pkg.clientsconfig["clientstask"] !="unkown"{
// 	g := pkg.ClientStartInput()
// }
