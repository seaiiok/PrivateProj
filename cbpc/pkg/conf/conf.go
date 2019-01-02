package conf

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

//getConf get user config file
func getConf(path string) (map[string]string, error) {
	c := make(map[string]string)
	f, err := os.Open(path)
	if err != nil {
		return c, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return c, err
		}

		s := strings.TrimSpace(string(b))

		expa := "#|/"
		expc := "="

		arri, reg := compile(expa, s)
		if len(arri) > 0 && arri[0] == 0 {
			continue
		}
		arri, reg = compile(expc, s)
		if len(arri) > 0 {
			arr := reg.Split(s, 2)
			c[strings.TrimSpace(arr[0])] = strings.TrimSpace(strings.TrimSpace(arr[1]))
		}
	}
	return c, err
}

func compile(expr string, s string) ([]int, *regexp.Regexp) {
	reg, _ := regexp.Compile(expr)
	m := reg.FindStringIndex(s)
	return m, reg
}
