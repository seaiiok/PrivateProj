package service

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
)

//功能追溯
//通常每个函数都需调用
func (p *MyProto) checkInit() MyProto {
	pc, _, _, _ := runtime.Caller(1)
	fname := runtime.FuncForPC(pc).Name()
	if p.Header.Error != nil {
		p.Header.DataTrace = append(p.Header.DataTrace, "Exist Error(Proto) before call function "+fname)
		return *p
	}
	p.Header.DataTrace = append(p.Header.DataTrace, fname+" -OK")
	return *p
}

//err队列 默认不超过100条
func (p *MyProto) appenderrors(err error){
	lenght:=len(p.Header.DataTrace)
	if lenght>=Proto_Errors {
		p.Header.DataTrace=p.Header.DataTrace[lenght-Proto_Errors:]	
	}
	p.Header.DataTrace = append(p.Header.DataTrace, err.Error())
}

//err的封装
func (p *MyProto) checkError(err error) error {
	if err != nil {
		p.appenderrors(err)
		p.Header.Error = err
		return err
	}
	return err
}

//客服端服务初始化
func (p *MyProto) clientInit() *MyProto {
	*p = p.checkInit()
	if p.Header.Error != nil {
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
func (p *MyProto) httpPost() *MyProto {

	*p = p.checkInit()
	if p.Header.Error != nil {
		return p
	}

	//创建请求
	postReq, err := http.NewRequest(p.Header.Method,
		p.Header.Url,
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

//struct2reader
func (p *MyProto) struct2reader() (io.Reader) {
	b, err := json.Marshal(p)
    p.checkError(err)
	return bytes.NewReader(b)
}

//reader2struct
func (p *MyProto) reader2struct(r io.Reader) {
	b, err := ioutil.ReadAll(r)
    p.checkError(err)
	json.Unmarshal(b, &p)
}

//struct2arraybytes
func (p *MyProto) struct2arraybytes() ([]byte) {
	b, err := json.Marshal(p)
    p.checkError(err)
	return b
}

//map2struct
func map2struct(m []map[string]string,s interface{}) (res []interface{}) {
	for i := 0; i < len(m); i++ {
		jsonStr, _ := json.Marshal(m[i])
		json.Unmarshal(jsonStr, &s)
		res=append(res,s)
	}
return
}

