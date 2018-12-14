package pkg

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"runtime"
)
func NewProto() *Proto {
	_header:=header{	
		ProcessTrace:make([]string,0),
		HeadMsg:make(map[string]string,0),
	}
	_bodymsg:=make([]map[string]string,0)
	_body:=body{
		BodyData:_bodymsg,
		// BodyMsg:make(map[int]map[string]string,0),
	}

return &Proto{}
}
//功能追溯
//通常每个函数都需调用
func (p *Proto) functraceInit() Proto {
	pc, _, _, _ := runtime.Caller(1)
	fname := runtime.FuncForPC(pc).Name()
	if p.header.Error != nil {
		p.header.DataTrace = append(p.header.DataTrace, "Exist Error(Proto) before call function "+fname)
		return *p
	}
	p.header.DataTrace = append(p.header.DataTrace, fname+" -OK")
	return *p
}

//err队列 默认不超过100条
func (p *Proto) appenderrors(err error) {
	lenght := len(p.header.DataTrace)
	if lenght >= Proto_Errors {
		p.header.DataTrace = p.header.DataTrace[lenght-Proto_Errors:]
	}
	p.header.DataTrace = append(p.header.DataTrace, err.Error())
}

//err的封装
func (p *Proto) checkError(err error) error {
	if err != nil {
		p.appenderrors(err)
		p.header.Error = err
		return err
	}
	return err
}

//struct2reader
func (p *Proto) struct2reader() io.Reader {
	b, err := json.Marshal(p)
	p.checkError(err)
	return bytes.NewReader(b)
}

//reader2struct
func (p *Proto) reader2struct(r io.Reader) {
	b, err := ioutil.ReadAll(r)
	p.checkError(err)
	json.Unmarshal(b, &p)
}

//struct2arraybytes
func (p *Proto) struct2arraybytes() []byte {
	b, err := json.Marshal(p)
	p.checkError(err)
	return b
}

//map2struct
func map2struct(m []map[string]string, s interface{}) (res []interface{}) {
	for i := 0; i < len(m); i++ {
		jsonStr, _ := json.Marshal(m[i])
		json.Unmarshal(jsonStr, &s)
		res = append(res, s)
	}
	return
}
