package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// GetAllFile1th("D:\\FTP\\offline\\ABC")
	// fileslist=make([]string,0)
	// err:=GetAllFile("D:\\FTP\\offline\\ABC")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for i := 0; i < len(fileslist); i++ {
	// 	fmt.Println(fileslist[i])
	// }
	path:="D:\\FTP\\offline\\ABC"
	f:=GetAllFiles(path)
	fl,_:=f(path)
	for i := 0; i < len(fl); i++ {
		fmt.Println(fl[i])
	}
}

// func GetAllFile(pathname string) ([]string,error) {
// 	rd, err := ioutil.ReadDir(pathname)
// 	for _, fi := range rd {
// 		if fi.IsDir() {
// 			GetAllFile(pathname +"\\"+ fi.Name(),fileslist)
// 		} else {
// 			fileslist=append(fileslist,pathname +"\\"+ fi.Name())
// 		}	
// 	}
// 	return
// }
type fc func(string)([]string,error)
func GetAllFiles(pathname string) (f fc) {
	fileslist:=make([]string,0)
	f=func (pathname string) ([]string,error) {
		rd, err := ioutil.ReadDir(pathname)
		for _, fi := range rd {
			if fi.IsDir() {
				f(pathname +"\\"+ fi.Name())
			} else {
				fileslist=append(fileslist,pathname +"\\"+ fi.Name())
			}	
		}
		return fileslist,err
	}
	return f
}
