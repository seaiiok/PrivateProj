package log

import(
"os"
"time"
"runtime"
"strings"
"strconv"
)

//日志功能
func Log(msg interface{}){
	var logmsg string
	filename:="./log/"+time.Now().Format("200601")+".log"
	_, file, line, _:= runtime.Caller(1)
	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND,0755)
	lines:=strconv.Itoa(line) 
	files:=strings.Replace(file,".go","",-1)

	switch msg.(type) {
	case string:
		logmsg=msg.(string)
		f.WriteString("---------------- "+time.Now().Format("2006-01-02 15:04:05")+" ----------------"+"\r\n"+"log: "+logmsg+"\r\n"+"\r\n"+"\r\n")
	case error:
		logmsg=msg.(error).Error()
		f.WriteString("---------------- "+time.Now().Format("2006-01-02 15:04:05")+" ----------------"+"\r\n"+"log: "+logmsg+"\r\n"+"loc: "+files+"  line:"+lines+"\r\n"+"\r\n"+"\r\n")
	default :
		logmsg="me(log) error"
		f.WriteString("---------------- "+time.Now().Format("2006-01-02 15:04:05")+" ----------------"+"\r\n"+"log: "+logmsg+"\r\n"+"\r\n"+"\r\n")
	}

	
}


