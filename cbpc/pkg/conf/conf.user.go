package conf

import (
	"fmt"
	"os"
)

//Config public config
var Config map[string]string

func init() {
	GetConfig("ifixtools.conf")
}

//GetConfig get user config file
func GetConfig(path string) (err error) {
	Config = make(map[string]string, 0)
	Config, err = getConf(path)
	if err != nil {
		Config = make(map[string]string, 0)
	}
	return
}

//SetConf save config
func SetConf(path string, conf map[string]string) (err error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer f.Close()
	for k, v := range conf {
		_, err = f.WriteString(k + "=" + v + "\n")
		fmt.Println(err)
	}
	return err
}
