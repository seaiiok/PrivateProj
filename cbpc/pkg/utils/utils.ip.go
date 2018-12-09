package utils

import(
    "os"
	"net"
)

//指出生产网IP
func GetLocalIp() (Ip []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Exit(1)
	}
	for _, address := range addrs {
		
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				Ip=append(Ip,ipnet.IP.String())
			}
		}
    }
		return
}

//指出匹配IP
func GetLocalMatchIp(match string) (Ip string) {
	addrs, err := net.InterfaceAddrs()
	ip:=[]string{""}
	if err != nil {
		os.Exit(1)
	}
	for _, address := range addrs {
		
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip=append(ip,ipnet.IP.String())
			}
		}
    }
		Ip=RegExpArrayMatch(match,ip)
		return
}
