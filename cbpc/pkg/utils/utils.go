package utils

import "strings"

//ArrayDiff array diff --keep scrarr
func ArrayDiff(srcarr []string, desarr []string) (res []string) {
	if len(desarr) == 0 {
		return srcarr
	}
	diff := false
	for i := 0; i < len(srcarr); i++ {
		diff = false
		for j := 0; j < len(desarr); j++ {
			if strTrim(srcarr[i]) == strTrim(desarr[j]) || len(strTrim(srcarr[i])) == 0 {
				break
			} else if j == len(desarr)-1 {
				diff = true
			}
		}
		if diff == true {
			res = append(res, srcarr[i])
		}
	}
	return
}

func strTrim(str string) (res string) {
	res = strings.TrimSpace(str)
	return
}
