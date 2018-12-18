package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
)

// func (p *Proto)GetFilesPath() ([]string, error) {
// ex,err:=pathExists(p.body.MyConf[Proto_SQL_clientsfilespath])
// }

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile(pathname + fi.Name() + "\\")
		} else {
			fmt.Println(fi.Name())
		}
	}
	return err
}
