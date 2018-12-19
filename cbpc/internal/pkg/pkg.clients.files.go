package pkg

import (
	"ifix.cbpc/cbpc/pkg/files"
)

func GetFileLines(path string) []string {	
	return files.ReadZipLines(path)
}

func GetFilePaths(root string,inp []int) (res []string) {	
	res=make([]string,0)
	f:=files.GetAllFiles()
	lp,err:=f(root)
	if err != nil {
		return 
	}
	res=files.GetNeedFiles(lp,inp)
	return
}