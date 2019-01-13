package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	a := "130146940179 PK77850007          1     1      48509     1     1 80 D2"
	b := "130147298019 PK78137486          2                              84 SH"
	c := []string{a, b}

	res := GetRowsData(c)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func GetRowsData(row []string) [][]string {
	res := make([][]string, 0)
	GoodSum := 0
	BadSum := 0

	for i := 0; i < len(row); i++ {
		re := make([]string, 0)
		rowarr := strings.Split(row[i], " ")
		lenrow := len(rowarr)
		if lenrow >= 4 {
			StdNo := rowarr[0]
			PKNo := rowarr[1]
			Coder := rowarr[lenrow-2]
			Dout := rowarr[lenrow-1]
			if len(StdNo) >= 12 && len(PKNo) >= 10 && len(Coder) == 2 && len(Dout) == 2 {
				if Coder == "80" {
					GoodSum++
				} else {
					BadSum++
				}
				re = append(re, strconv.Itoa(GoodSum), strconv.Itoa(BadSum), StdNo, PKNo, Coder, Dout)
			}
		}
		if len(re) == 6 {
			res = append(res, re)
		}
	}
	return res
}
