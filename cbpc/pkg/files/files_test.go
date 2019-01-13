package files

import (
	"fmt"
	"testing"
)

func TestGetLine(t *testing.T) {
	// type fc func(string) ([]string, error)
	// func GetAllFiles() (f fc) {

	f := GetAllFiles()
	path, err := f("D:/FTP/online/ABC")
	t.Log(err)
	inp := "5"
	res1 := GetNeedFilesInp(path, inp)
	// t.Log(res1)
	res2 := GetNeedFilesMatch(res1, ".ZIP|.zip")
	for _, v := range res2 {
		fmt.Println(v)
	}


}
