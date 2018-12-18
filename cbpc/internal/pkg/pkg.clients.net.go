package pkg

import (
	"crypto/tls"
	"errors"
	"net/http"
)

var client *http.Client

// ClientsHTTPStart 客服端
func ClientsHTTPStart() {
	p := clientsHTTPInit()
	go p.clientsiFixToolsController()
}

//客服端服务初始化
func clientsHTTPInit() *Proto {
	p := NewProto()
	p.HeadMsg["method"] = "POST"
	p.HeadMsg["url"] = "https://" + config["serveraddr"] + ":" + config["serverport"] + "/ifixtools"
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
	if p.header.Error != nil {
		p.ProcessTrace()
		return p
	}

	//创建请求
	postReq, err := http.NewRequest(p.header.HeadMsg["method"],
		p.header.HeadMsg["url"],
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
	p.checkError(err)
	if err != nil {
		return p
	}
	p.reader2struct(resp.Body)
	if resp.StatusCode != 200 {
		p.header.Error = errors.New(resp.Status)
	}
	defer resp.Body.Close()
	return p
}
