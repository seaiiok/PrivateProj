package protocol

import (
	"runtime"
)

const (
	sum = 10
)

//NewProto init user proto
func NewProto() *Protocol {
	return &Protocol{
		Error:        nil,
		ProcessTrace: make([]string, 0),
		MeConf:       make(map[string]string, 0),
		Data:         make([]string, 0),
		Datas:        make([][]string, 0),
	}
}

//SetProcessTrace init user proto
func (p *Protocol) SetProcessTrace(str ...string) {
	l := len(p.ProcessTrace)
	if l >= sum {
		p.ProcessTrace = p.ProcessTrace[l-sum:]
	}

	s := ""
	for _, v := range str {
		s += v
	}
	if s != "" {
		p.ProcessTrace = append(p.ProcessTrace, s)
		return
	}
	pc, _, _, _ := runtime.Caller(1)
	fname := runtime.FuncForPC(pc).Name()
	p.ProcessTrace = append(p.ProcessTrace, fname)
}

// //SetError big error that donot go user proto/clear
// func (p *protocol) SetError(err error) {
// 	p.Error = err
// 	p.setError()
// }

// func (p *protocol) setError() {
// 	p.Data = make([]string, 0)
// 	p.Datas = make([][]string, 0)
// }

// //Struct2Reader ...
// func (p *protocol) Struct2Reader() io.Reader {
// 	b, _ := json.Marshal(p)
// 	return bytes.NewReader(b)
// }

// //Reader2Struct ...
// func (p *protocol) Reader2Struct(r io.Reader) {
// 	b, _ := ioutil.ReadAll(r)
// 	json.Unmarshal(b, &p)
// }
