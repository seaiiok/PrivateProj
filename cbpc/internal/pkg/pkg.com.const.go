package pkg

const (
	Proto_Errors = iota
	Proto_Id_init
	Proto_Id_Rdy
	Proto_Id_Nil
	Proto_Id_Err
	Proto_Id_Start
	Proto_Id_Stop
	Proto_Id_Restart
)

const (
	Proto_Cmd                      = "Ladygaga"
	Proto_Cmd_Resp                 = "响应"
	Proto_Cmd_Resp_Yes             = "YES"
	Proto_Cmd_Resp_No              = "NO"
	Proto_Cmd_Resp_Again           = "重来"
	Proto_Cmd_Req_Ready            = "准备"
	Proto_Cmd_Req_KeyCols          = "唯一键"
	Proto_Cmd_Req_InsertRows       = "存储"
	Proto_Cmd_ServerDriver         = "金刚芭比-小哪吒"
	Proto_Map_ServerConn           = "是他,是他,是他,就是他,我们的朋友小哪吒"
	Proto_Map_ServerTableName      = "是他,就是他,是他,就是他,少年英雄小哪吒"
	Proto_Map_ServerTableCol       = "上天他比天要高,下海他比海更大啊~啊~"
	Proto_Map_ServerSqlCreateTable = "智斗妖魔勇降鬼怪,少年英雄就是小哪吒"
	Proto_Map_ServerSqlCol         = "有时他很聪明,有时他也犯傻,他的个头跟我一般高"
	Proto_Map_ServerSqlExistTable  = "有时他很努力,有时他也贪玩,他的年纪和我一般大"
	Proto_Map_ClientsDriver        = "啦啦啦啦~啦啦啦啦~1"
	Proto_Map_ClientsConn          = "啦啦啦啦~啦啦啦啦~2"
	Proto_Map_ClientsTableName     = "啦啦啦啦~啦啦啦啦~3"
	Proto_Map_ClientsTableCol      = "啦啦啦啦~啦啦啦啦~4"
	Proto_Map_ClientsSqlCol        = "啦啦啦啦~啦啦啦啦~5"
	Proto_Map_ClientsSqlExistTable = "啦啦啦啦~啦啦啦啦~6"
	Proto_Map_ClientsSql           = "啦啦啦啦~啦啦啦啦~7"
	Proto_Map_ClientsTask          = "啦啦啦啦~啦啦啦啦~8"
)

//配置文件Map
const (
	Proto_Conf_serverdbkeycol            = "serverdbkeycol"
	Proto_Conf_clientsd3weightselectrows = "clientsd3weightselectrows"
	Proto_Conf_clientsd3weightinsertrows = "clientsd3weightinsertrows"
)
