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
	Proto_Cmd_Restart              = "重来"
	Proto_Cmd_Again                = "再来"
	Proto_Cmd_1th                  = "Proto_Cmd_1th"
	Proto_Cmd_2th                  = "Proto_Cmd_2th"
	Proto_Cmd_3th                  = "Proto_Cmd_3th"
	Proto_Cmd_4th                  = "Proto_Cmd_4th"
	Proto_Cmd_5th                  = "Proto_Cmd_5th"
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

//严格的
const (
	Proto_Conf_ProcessTraceSum = 5
	Proto_Conf_Task_d3weight = "d3weight"
	Proto_Conf_Task_jwonlinefile = "jwonlinefile"
	Proto_Conf_Task_jwofflinedb = "jwofflinedb"
	Proto_Conf_Task_jfofflinefile = "jfofflinefile"
	Proto_Conf_Task_jfofflinedb = "jfofflinedb"
)


//配置文件Map
const (
	Proto_Conf_serverdbdriver            = "serverdbdriver"
	Proto_Conf_serverdbconn              = "serverdbconn"
	Proto_Conf_serverd3weightkeycol      = "serverd3weightkeycol"
	Proto_Conf_serverd3weightinsertrows  = "serverd3weightinsertrows"
	Proto_Conf_clientsd3weightdbdriver   = "clientsd3weightdbdriver"
	Proto_Conf_clientsd3weightdbconn     = "clientsd3weightdbconn"
	Proto_Conf_clientsd3weightdbkeycol   = "clientsd3weightdbkeycol"
	Proto_Conf_clientsd3weightkeycol     = "clientsd3weightkeycol"
	Proto_Conf_clientsd3weightselectrows = "clientsd3weightselectrows"
	Proto_Conf_clientstask                 = "clientstask"
)
