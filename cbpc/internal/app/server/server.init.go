package server

import (
	"fmt"

	"ifix.cbpc/cbpc/internal/pkg"
)

func serverInit() {
	pkg.ServerConfigInit("")
	pkg.ServerDBInit()

	path := "D://FTP//offline//ABC//2017//A002//KNR04886_012_10.zip"
	inp := []int{6, 7}
	list := pkg.GetFileLines(path, inp)
	fmt.Println(list)
}
