package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "D://FTP//offline//ABC//2017//A002//KNR04886_012_1.zip"
	f := ReadZipLines(path)
	fmt.Println(f)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadZipLines(path string) (res []string) {
	res = make([]string, 0)
	b,_:=PathExists(path)
	if b !=true {
		return res
	}
	rc, err := zip.OpenReader(path)

	if err != nil {
		defer rc.Close()
	}

	for _, _file := range rc.File {
		f, _ := _file.Open()
		rd := bufio.NewReader(f)
		for {
			line, err := rd.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			res = append(res, line)
		}
		defer f.Close()
	}
	return
}
