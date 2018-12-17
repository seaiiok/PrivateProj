package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"ifix.cbpc/cbpc/pkg/print"
)

//NewProto .
func NewProto() *Proto {
	return &Proto{
		header: header{
			ProcessTrace: make([]string, 0),
			HeadMsg:      make(map[string]string, 0),
		},
		body: body{
			MyConf:    make(map[string]string, 0),
			BodyData:  make([]string, 0),
			BodyDatas: make([][]string, 0),
		},
	}
}

//ProcessTrace 功能追溯
//通常每个函数都需调用
func (p *Proto) ProcessTrace(str ...string) *Proto {

	l := len(p.header.ProcessTrace)
	if l >= Proto_Conf_ProcessTraceSum {
		p.header.ProcessTrace = p.header.ProcessTrace[l-Proto_Conf_ProcessTraceSum:]
	}

	s := ""
	for _, v := range str {
		s += v
	}
	if s != "" {
		p.header.ProcessTrace = append(p.header.ProcessTrace, s)
		return p
	}
	pc, _, _, _ := runtime.Caller(1)
	fname := runtime.FuncForPC(pc).Name()
	p.header.ProcessTrace = append(p.header.ProcessTrace, fname)
	return p
}

//checkError .
func (p *Proto) checkError(err error) {
	p.header.Error = err
}

//struct2reader
func (p *Proto) struct2reader() io.Reader {
	b, _ := json.Marshal(p)
	return bytes.NewReader(b)
}

//reader2struct
func (p *Proto) reader2struct(r io.Reader) {
	b, _ := ioutil.ReadAll(r)
	json.Unmarshal(b, &p)
}

//struct2arraybytes
func (p *Proto) struct2arraybytes() []byte {
	b, _ := json.Marshal(p)	
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

//数组差异集,方法过时
func ArrayDiff(srcarr []string, desarr []string) (res []string) {
	if len(desarr) == 0 {
		return srcarr
	}
	diff := false
	for i := 0; i < len(srcarr); i++ {
		diff = false
		for j := 0; j < len(desarr); j++ {
			if strTrim(srcarr[i]) == strTrim(desarr[j]) || len(strTrim(srcarr[i])) == 0 {
				break
			} else if j == len(desarr)-1 {
				diff = true
			}
		}
		if diff == true {
			res = append(res, srcarr[i])
		}
	}
	return
}

func strTrim(str string) (res string) {
	res = strings.TrimSpace(str)
	return
}

// same /pkg
func (p *Proto)sqlStringMakeDTimes(ostr string) {
	p.body.MyConf[Proto_SQL_clientsdbcol]="select F_ID from T_Standard where F_CNTime > '2017-01-01 00:00:00'"
	p.body.MyConf[Proto_SQL_serverdbcol]="select F_ID from iFixsvr_D3Weight where F_CNTime > '2017-01-01 00:00:00'"
	// _ostr := strings.Split(ostr, "\\")
	// if len(_ostr) != 2 {
	// 	p.body.MyConf[Proto_SQL_clientsdbcol]=  ostr
	// }
	// p.body.MyConf[Proto_SQL_clientsdbcol]= _ostr[0] + sqlStringMakeDTime(_ostr[1])	
}

// same /pkg
func sqlStringMakeDTime(s string) string {
	c := regexp.MustCompile("[0-9]+")
	p := c.FindString(s)

	i, _ := strconv.Atoi(p)
	if i > 3650 {
		i = 3650
	}
	if i < 0 {
		i = 0
	}
	i = -24 * i

	d, _ := time.ParseDuration(strconv.Itoa(i) + "h")
	dt := time.Now().Add(d)
	return "'" + dt.Format("2006-01-02 15:04:05") + "'"

}

func (p *Proto)sqlStringServerMakeRows(ostr string) {
	p.body.MyConf[Proto_SQL_serverdbrows]= ostr
}

func (p *Proto)sqlStringClientsMakeRows(ostr string, col []string) {
	_ostr := strings.Split(ostr, "\\")
	if len(_ostr) != 2 {
		p.body.MyConf[Proto_SQL_clientsdbrows]= ostr
	}
	fmt.Println(ostr)
	fmt.Println(_ostr[0])
	fmt.Println(sqlStringMakeRow(col))
	p.body.MyConf[Proto_SQL_clientsdbrows]= _ostr[0] + sqlStringMakeRow(col)
}

func sqlStringMakeRow(ostr []string) (nstr string) {

	for i := 0; i < len(ostr); i++ {
		if i > 0 {
			nstr += ",'" + ostr[i] + "'"
		} else {
			nstr = "'" + ostr[i] + "'"
		}
	}
	return " (" + nstr + ")"
}

func ClientStartInput() string {
	print.Line(print.Ok, "iFix工具已启动,请选择一个任务!")
	print.Line(print.Ok, "1:3#门车辆称重")
	print.Line(print.Ok, "2:在线清数")
	print.Line(print.Ok, "3:离线检测")
	print.Line(print.Ok, "4:检封工房离线数据(数据包)")
	print.Line(print.Ok, "5:检封工房离线数据(数据库)")
	task := 0
	fmt.Scanln(&task)
	switch task {
	case 1:
		return Proto_Conf_Task_d3weight
	case 2:
		return Proto_Conf_Task_jwonlinefile
	case 3:
		return Proto_Conf_Task_jwofflinedb
	case 4:
		return Proto_Conf_Task_jfofflinefile
	case 5:
		return Proto_Conf_Task_jfofflinedb
	default:
		return ClientStartInput()
	}
}
