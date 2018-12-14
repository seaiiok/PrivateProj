package pkg

import (
	"crypto/tls"
	"net/http"
	"fmt"
)

func ClientsHttpStart() {
	p := new(Proto)
	p.HeadMsg["method"]= "POST"
	p.HeadMsg["url"]= "https://192.168.100.38:494/test"
	p.clientInit()
	p1:=p.httpPost()
	fmt.Println(p1)
}

//客服端服务初始化
func (p *Proto) clientInit() *Proto {
	*p = p.checkInit()
	if p.header.Error != nil {
		return p
	}
	//忽略证书校验
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
	return p
}

//WebService post请求
//内部
func (p *Proto) httpPost() *Proto {

	*p = p.checkInit()
	if p.header.Error != nil {
		return p
	}

	//创建请求
	postReq, err := http.NewRequest(p.header.Method,
		p.header.URL,
		p.struct2reader(),
	)
	p.checkError(err)
	if err != nil {
		return p
	}

	//增加header
	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")
	//执行请求
	resp, err := client.Do(postReq)

	err = p.checkError(err)
	if err != nil {
		return p
	}

	p.reader2struct(resp.Body)
	defer resp.Body.Close()
	return p
}
