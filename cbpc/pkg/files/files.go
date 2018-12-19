package files

import (
	"archive/zip"
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type fc func(string) ([]string, error)

func GetAllFiles() (f fc) {
	fileslist := make([]string, 0)
	f = func(pathname string) ([]string, error) {
		rd, err := ioutil.ReadDir(pathname)
		for _, fi := range rd {
			if fi.IsDir() {
				f(pathname + "\\" + fi.Name())
			} else {
				fileslist = append(fileslist, pathname+"\\"+fi.Name())
			}
		}
		return fileslist, err
	}
	return f
}

func GetNeedFiles(path []string, inp []int) (res []string) {
	for i := 0; i < len(path); i++ {
		paths := strings.Split(path[i], "\\")
		for j := 0; j < len(inp); j++ {
			if len(paths) == inp[j] {
				res = append(res, path[i])
			}
		}
	}
	return
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
	b, _ := PathExists(path)
	if b != true {
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
