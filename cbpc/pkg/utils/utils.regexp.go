package utils

import (
	"regexp"
)

//匹配
func RegExpMatch(exp string,s string) bool {

	b,_:=regexp.MatchString(exp,s)
     return b
}

//匹配
func RegExpArrayMatch(exp string,s []string) string {
	for i := 0; i < len(s); i++ {
		b:=RegExpMatch(exp,s[i])
		if b == true {
			return s[i]	
		}
	}
return ""
}

