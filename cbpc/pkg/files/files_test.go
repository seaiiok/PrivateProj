package files

import (
	"testing"
)

func TestGetLine(t *testing.T) {
	// type fc func(string) ([]string, error)
	// func GetAllFiles() (f fc) {

	f := GetAllFiles()
	path, err := f("D:/FTP/online/ABC")
	t.Log(err)
	inp := []int{5}
	res := GetNeedFiles(path, inp)
	t.Log(res)
	rows:=ReadZipLines(res[0])
	t.Log(rows)
}
