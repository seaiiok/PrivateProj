package conf

import(
	"testing"
)

func TestSetConfKV(t *testing.T){

		path := "./ifixtools.conf"
	m:=make(map[string]string,0)
	m["a"]="A"
	m["b"]="B"

	SetConfKV(path,m)
}