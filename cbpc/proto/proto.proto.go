package proto

import (
	"runtime"
)

const (
	sum = 10
)

//NewProto init user proto
func NewProto() *Proto {
	return &Proto{
		Err: Err{
			Err:          nil,
			processTrace: make([]string, 0),
		},
		Config: Config{
			Config: make(map[string]string, 0),
		},
		Sql: Sql{
			DatabaseDriver: "",
			DatabaseSource: "",
			Insert:         "",
			Query:          "",
			Args:           make([]string, 0),
			Data:           make([][]string, 0),
		},
		Device: Device{
			DeviceName:   "",
			DeviceIP:     "",
			DeviceTask:   "",
			DeviceRouter: "",
		},
	}
}

//SetProcessTrace add ProcessTrace
func (p *Proto) SetProcessTrace(str ...string) {
	l := len(p.Err.processTrace)
	if l >= sum {
		p.Err.processTrace = p.Err.processTrace[l-sum:]
	}

	s := ""
	for _, v := range str {
		s += v
	}
	if s != "" {
		p.Err.processTrace = append(p.Err.processTrace, s)
		return
	}
	pc, _, _, _ := runtime.Caller(1)
	fname := runtime.FuncForPC(pc).Name()
	p.Err.processTrace = append(p.Err.processTrace, fname)
}

//GetProcessTrace add ProcessTrace
func (p *Proto) GetProcessTrace() []string {
	return p.Err.processTrace
}

//GetFulQuery return query
func (p *Proto) GetFulQuery() string {
	fullquery := p.Sql.Query
	fullparams := ""
	if len(p.Sql.Args) <= 0 {
		return fullquery
	}
	for i := 0; i < len(p.Sql.Args); i++ {
		fullparams += "'" + p.Sql.Args[i] + "',"
	}
	fullparams = fullparams[:len(fullparams)-1]
	return p.Sql.Query + " (" + fullparams + ")"
}

//GetQuery return query
func (p *Proto) GetQuery() string {
	return p.Sql.Query
}

//SetQuery set query
func (p *Proto) SetQuery(query string) {
	p.Sql.Query = query
}

//SetArgs set args
func (p *Proto) SetArgs(args []string) {
	p.Sql.Args = make([]string, 0)
	p.Sql.Args = args
}

//GetArgs return args
func (p *Proto) GetArgs() []string {
	return p.Sql.Args
}

//SetData set data
func (p *Proto) SetData(data [][]string) {
	p.Sql.Data = make([][]string, 0)
	p.Sql.Data = data
}

//GetData return data
func (p *Proto) GetData() [][]string {
	return p.Sql.Data
}

//SetError set data
func (p *Proto) SetError(err error) {
	if p.Err.Err != nil {
		return
	}
	p.Err.Err=err
}

//SetError set data
func (p *Proto) SetRouter(router string) {
	if p.Err.Err != nil {
		p.DeviceRouter="routerSetObjectRestart"
		return
	}
	p.DeviceRouter=router
}


