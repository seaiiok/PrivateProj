package service

//初始化报文
//用于ClientsInit
func ClientsProtoInit () (*MyProto) {
	//MyProto.Header
	//报文头部
	newp:=new(MyProto)
	newp.Header.Id = Proto_Id_init
	newp.Header.Method = "POST"
	newp.Header.Identify=false
	newp.Header.Url = "https://127.0.0.1:8080/test"
	newp.Header.Error=nil
	newp.Header.DataTrace=make([]string,0,Proto_Errors)
	newp.Header.HeadMsg = make(map[string]string, 0)

	//MyProto.Body
	//报文内容
	newp.Body.BodyMsg = make(map[int]map[string]string, 0)
	return newp
}

//初始化报文
//用于服务器响应初始化报文
func (p *MyProto)ServerProtoInit () {
	//MyProto.Header
	//报文头部
	p.Header.Id = Proto_Id_init
	p.Header.Method = "POST"
	p.Header.Identify=false
	p.Header.Url = "https://127.0.0.1:8080/test"
	p.Header.Error=nil
	p.Header.DataTrace=make([]string,0,Proto_Errors)
	p.Header.HeadMsg = make(map[string]string, 0)
	p.Header.HeadMsg=getHeadMsg()
	//MyProto.Body
	//报文内容
	p.Body.BodyMsg = make(map[int]map[string]string, 0)

}

func getHeadMsg () map[string]string {
words:=make(map[string]string, 0)
words[Proto_Map_ServerDriver]="odbc"
words[Proto_Map_ServerConn]="server=192.168.1.8;user id =sa;password=seaii1949!;port=1433;database=iFixsvr"
words[Proto_Map_ServerTableName]=""
words[Proto_Map_ServerTableCol]=""
words[Proto_Map_ServerSqlCreateTable]=""
words[Proto_Map_ServerSqlCol]=""
words[Proto_Map_ServerSqlExistTable]=""
words[Proto_Map_ClientDriver]="odbc"
words[Proto_Map_ClientConn]=""
words[Proto_Map_ClientTableName]=""
words[Proto_Map_ClientTableCol]=""
words[Proto_Map_ClientSqlCol]=""
words[Proto_Map_ClientSqlExistTable]=""
words[Proto_Map_ClientSql]=""
words[Proto_Map_ClientTask]=""
return words
}

/** HeadMsg 常用字典
did 
设备唯一标识号  数采合同可查

taskname 
任务名

serverdbdriver  
服务器数据库使用的驱动名

serverdbconnstr  
服务器数据库连接字



*/

