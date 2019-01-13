package protocol

import (
	"runtime"
)

const (
	sum           = 10
	routerRestart = "/routergetobjectrestart"
)

//NewProto init user proto
func NewProto() *Proto {
	return &Proto{
		Err: Err{
			Error:        false,
			ProcessTrace: make([]string, 0),
		},
		SQL: SQL{
			DatabaseDriver: "",
			DatabaseSource: "",
			InsertSQL:      "",
			QuerySQL:       "",
			Args:           make([]string, 0),
			Data:           make([][]string, 0),
		},
		Device: Device{
			DeviceName:   "",
			DeviceIP:     "",
			DeviceTask:   "",
			DeviceRouter: "",
			DeviceStatusCode:0,
		},
	}
}

//SetProcessTrace add ProcessTrace
func (p *Proto) SetProcessTrace(str ...string) {
	l := len(p.Err.ProcessTrace)
	if l >= sum {
		p.Err.ProcessTrace = p.Err.ProcessTrace[l-sum:]
	}

	s := ""
	for _, v := range str {
		s += v
	}
	if s != "" {
		p.Err.ProcessTrace = append(p.Err.ProcessTrace, s)
		return
	}
	pc, _, _, _ := runtime.Caller(1)
	fname := runtime.FuncForPC(pc).Name()
	p.Err.ProcessTrace = append(p.Err.ProcessTrace, fname)
}


//SetArgs set args
func (p *Proto) SetArgs(args []string) {
	p.SQL.Args = make([]string, 0)
	p.SQL.Args = args
}

//SetData set data
func (p *Proto) SetData(data [][]string) {
	p.SQL.Data = make([][]string, 0)
	p.SQL.Data = data
}

//GetData return data
func (p *Proto) GetData() [][]string {
	return p.SQL.Data
}

//SetRouter set data
func (p *Proto) SetRouter(router string) {
	if p.Err.Error != false {
		p.DeviceRouter = routerRestart
		return
	}
	p.DeviceRouter = router
}


//SetError set err
func (p *Proto) SetError(err bool) {
	if p.Err.Error != false {
		return
	}
	p.Err.Error = err
}

//SetErrorNil set err nil
func (p *Proto) SetErrorNil() {
	p.Err.Error = false
}