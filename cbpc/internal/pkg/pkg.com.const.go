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
	Proto_Cmd_6th                  = "Proto_Cmd_6th"
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
	Proto_Conf_ProcessTraceSum = 10
	Proto_Conf_ProcessSleep    = 5
)

//严格的
const (
	Proto_Conf_Task_d3weight      = "d3weight"
	Proto_Conf_Task_jwonlinefile  = "jwonlinefile"
	Proto_Conf_Task_jwofflinedb   = "jwofflinedb"
	Proto_Conf_Task_jfofflinefile = "jfofflinefile"
	Proto_Conf_Task_jfofflinedb   = "jfofflinedb"
	Proto_Conf_Task_unknown       = "unknown"
	Proto_Conf_clientstask        = "clientstask"
)

//严格的
const (
	Proto_SQL_serverdbdriver      = "serverdbdriver"
	Proto_SQL_serverdbconn        = "serverdbconn"
	Proto_SQL_serverdbcol         = "serverdbcol"
	Proto_SQL_serverdbrows        = "serverdbrows"
	Proto_SQL_clientsdbdriver     = "clientsdbdriver"
	Proto_SQL_clientsdbconn       = "clientsdbconn"
	Proto_SQL_clientsdbcol        = "clientsdbcol"
	Proto_SQL_clientsdbrows       = "clientsdbrows"
	Proto_SQL_clientsfilespath    = "clientsfilespath"
	Proto_SQL_clientsfilecreatedt = "clientsfilecreatedt"
)

//配置文件Map
const (
	Proto_Conf_serverdbdriver = "serverdbdriver"
	Proto_Conf_serverdbconn   = "serverdbconn"

	Proto_Conf_serverdbcol4d3weight     = "serverdbcol4d3weight"
	Proto_Conf_serverdbrows4d3weight    = "serverdbrows4d3weight"
	Proto_Conf_clientsdbdriver4d3weight = "clientsdbdriver4d3weight"
	Proto_Conf_clientsdbconn4d3weight   = "clientsdbconn4d3weight"
	Proto_Conf_clientsdbcol4d3weight    = "clientsdbcol4d3weight"
	Proto_Conf_clientsdbrows4d3weight   = "clientsdbrows4d3weight"

	Proto_Conf_serverdbcol4jwonlinefile     = "serverdbcol4jwonlinefile"
	Proto_Conf_serverdbrows4jwonlinefile    = "serverdbrows4jwonlinefile"
	Proto_Conf_clientsdbdriver4jwonlinefile = "clientsdbdriver4jwonlinefile"
	Proto_Conf_clientsdbconn4jwonlinefile   = "clientsdbconn4jwonlinefile"
	Proto_Conf_clientsdbcol4jwonlinefile    = "clientsdbcol4jwonlinefile"
	Proto_Conf_clientsdbrows4jwonlinefile   = "clientsdbrows4jwonlinefile"

	Proto_Conf_serverdbcol4jwofflinedb     = "serverdbcol4jwofflinedb"
	Proto_Conf_serverdbrows4jwofflinedb    = "serverdbrows4jwofflinedb"
	Proto_Conf_clientsdbdriver4jwofflinedb = "clientsdbdriver4jwofflinedb"
	Proto_Conf_clientsdbconn4jwofflinedb   = "clientsdbconn4jwofflinedb"
	Proto_Conf_clientsdbcol4jwofflinedb    = "clientsdbcol4jwofflinedb"
	Proto_Conf_clientsdbrows4jwofflinedb   = "clientsdbrows4jwofflinedb"

	Proto_Conf_serverdbcol4jfofflinefile       = "serverdbcol4jfofflinefile"
	Proto_Conf_serverdbrows4jfofflinefile      = "serverdbrows4jfofflinefile"
	Proto_Conf_clientsfilespath4jfofflinefile  = "clientsfilespath4jfofflinefile"
	Proto_Conf_clientsfilecreatedt4jfofflinefile = "clientsfilecreatedt4jfofflinefile"

	Proto_Conf_serverdbcol4jfofflinedb     = "serverdbcol4jfofflinedb"
	Proto_Conf_serverdbrows4jfofflinedb    = "serverdbrows4jfofflinedb"
	Proto_Conf_clientsdbdriver4jfofflinedb = "clientsdbdriver4jfofflinedb"
	Proto_Conf_clientsdbconn4jfofflinedb   = "clientsdbconn4jfofflinedb"
	Proto_Conf_clientsdbcol4jfofflinedb    = "clientsdbcol4jfofflinedb"
	Proto_Conf_clientsdbrows4jfofflinedb   = "clientsdbrows4jfofflinedb"
)
