package producers

import (
	"errors"
	"strings"
)

func getFulQuery(querysql string, args []string) (string, error) {
	fullparams := ""
	err := errors.New("error sql query string")
	if len(args) <= 0 {
		tempsql := strings.Split(querysql, "AND")
		for k, v := range tempsql {
			if len(tempsql) == 2 && k == 0 {
				return v, err
			}
		}
		return querysql, err
	}
	for i := 0; i < len(args); i++ {
		fullparams += "'" + args[i] + "',"
	}
	fullparams = fullparams[:len(fullparams)-1]
	return querysql + " (" + fullparams + ")", nil
}
