package utils

import(
	"testing"
)

func TestGetLocalIp(t *testing.T){
	ip:=GetLocalIp()
	for i := 0; i < len(ip); i++ {
		t.Log(ip[i])
	}
}

func TestGetLocalMatchIp(t *testing.T){
	ip:=GetLocalMatchIp("10.9")
		t.Log(ip)

}