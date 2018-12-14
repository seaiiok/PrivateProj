package pkg

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"runtime"
)

// //NewProto .
// func NewProto() *Proto {
// 	_header:=header{	
// 		ProcessTrace:make([]string,0),
// 		HeadMsg:make(map[string]string,0),
// 	}
// 	_body:=body{
// 		BodyData:make([]string,0),
// 		BodyDatas:make([][]string,0),
// 	}
// return &Proto{
// 	header:_header,
// 	body:_body,
// }
// }

//NewProto .
func NewProto() *Proto {
return &Proto{
	header:header{	
		ProcessTrace:make([]string,0),
		HeadMsg:make(map[string]string,0),
	},
	body:body{
		BodyData:make([]string,0),
		BodyDatas:make([][]string,0),
	},
}
}

//ProcessTrace 功能追溯
//通常每个函数都需调用
func (p *Proto) ProcessTrace() *Proto {
	pc, _, _, _ := runtime.Caller(1)
	fname := runtime.FuncForPC(pc).Name()
	if p.header.Error !=nil {
		p.header.ProcessTrace = append(p.header.ProcessTrace, "Exist Error(Proto) before call function "+fname)
		return p
	}
	p.header.ProcessTrace = append(p.header.ProcessTrace, fname+" -OK")
	return p
	
}

//checkError .
func (p *Proto) checkError(err error) {
	p.header.Error=err
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
