package pkg

import "errors"

//GetClientsConf config for all clients
func (p *Proto) GetServerConf() error {

	switch p.header.HeadMsg[Proto_Conf_clientstask] {
	case Proto_Conf_Task_d3weight:

		p.body.MyConf = make(map[string]string, 0)
		p.body.MyConf[Proto_SQL_serverdbdriver] = config[Proto_Conf_serverdbdriver]
		p.body.MyConf[Proto_SQL_serverdbconn] = config[Proto_Conf_serverdbconn]
		p.body.MyConf[Proto_SQL_serverdbcol] = config[Proto_Conf_serverdbcol4d3weight]
		p.body.MyConf[Proto_SQL_serverdbrows] = config[Proto_Conf_serverdbrows4d3weight]
		p.body.MyConf[Proto_SQL_clientsdbdriver] = config[Proto_Conf_clientsdbdriver4d3weight]
		p.body.MyConf[Proto_SQL_clientsdbconn] = config[Proto_Conf_clientsdbconn4d3weight]
		p.body.MyConf[Proto_SQL_clientsdbcol] = config[Proto_Conf_clientsdbcol4d3weight]
		p.body.MyConf[Proto_SQL_clientsdbrows] = config[Proto_Conf_clientsdbrows4d3weight]
	case Proto_Conf_Task_jwofflinedb:

		p.body.MyConf = make(map[string]string, 0)
		p.body.MyConf[Proto_SQL_serverdbdriver] = config[Proto_Conf_serverdbdriver]
		p.body.MyConf[Proto_SQL_serverdbconn] = config[Proto_Conf_serverdbconn]
		p.body.MyConf[Proto_SQL_serverdbcol] = config[Proto_Conf_serverdbcol4jwofflinedb]
		p.body.MyConf[Proto_SQL_serverdbrows] = config[Proto_Conf_serverdbrows4jwofflinedb]
		p.body.MyConf[Proto_SQL_clientsdbdriver] = config[Proto_Conf_clientsdbdriver4jwofflinedb]
		p.body.MyConf[Proto_SQL_clientsdbconn] = config[Proto_Conf_clientsdbconn4jwofflinedb]
		p.body.MyConf[Proto_SQL_clientsdbcol] = config[Proto_Conf_clientsdbcol4jwofflinedb]
		p.body.MyConf[Proto_SQL_clientsdbrows] = config[Proto_Conf_clientsdbrows4jwofflinedb]

	case Proto_Conf_Task_jwonlinefile:

		p.body.MyConf = make(map[string]string, 0)
		p.body.MyConf[Proto_SQL_serverdbdriver] = config[Proto_Conf_serverdbdriver]
		p.body.MyConf[Proto_SQL_serverdbconn] = config[Proto_Conf_serverdbconn]
		p.body.MyConf[Proto_SQL_serverdbcol] = config[Proto_Conf_serverdbcol4jwonlinefile]
		p.body.MyConf[Proto_SQL_serverdbrows] = config[Proto_Conf_serverdbrows4jwonlinefile]
		p.body.MyConf[Proto_SQL_clientsdbdriver] = config[Proto_Conf_clientsdbdriver4jwonlinefile]
		p.body.MyConf[Proto_SQL_clientsdbconn] = config[Proto_Conf_clientsdbconn4jwonlinefile]
		p.body.MyConf[Proto_SQL_clientsdbcol] = config[Proto_Conf_clientsdbcol4jwonlinefile]
		p.body.MyConf[Proto_SQL_clientsdbrows] = config[Proto_Conf_clientsdbrows4jwonlinefile]

	case Proto_Conf_Task_jfofflinedb:

		p.body.MyConf = make(map[string]string, 0)
		p.body.MyConf[Proto_SQL_serverdbdriver] = config[Proto_Conf_serverdbdriver]
		p.body.MyConf[Proto_SQL_serverdbconn] = config[Proto_Conf_serverdbconn]
		p.body.MyConf[Proto_SQL_serverdbcol] = config[Proto_Conf_serverdbcol4jfofflinedb]
		p.body.MyConf[Proto_SQL_serverdbrows] = config[Proto_Conf_serverdbrows4jfofflinedb]
		p.body.MyConf[Proto_SQL_clientsdbdriver] = config[Proto_Conf_clientsdbdriver4jfofflinedb]
		p.body.MyConf[Proto_SQL_clientsdbconn] = config[Proto_Conf_clientsdbconn4jfofflinedb]
		p.body.MyConf[Proto_SQL_clientsdbcol] = config[Proto_Conf_clientsdbcol4jfofflinedb]
		p.body.MyConf[Proto_SQL_clientsdbrows] = config[Proto_Conf_clientsdbrows4jfofflinedb]

	case Proto_Conf_Task_jfofflinefile:

		p.body.MyConf = make(map[string]string, 0)
		p.body.MyConf[Proto_SQL_serverdbdriver] = config[Proto_Conf_serverdbdriver]
		p.body.MyConf[Proto_SQL_serverdbconn] = config[Proto_Conf_serverdbconn]
		p.body.MyConf[Proto_SQL_serverdbcol] = config[Proto_Conf_serverdbcol4jfofflinefile]
		p.body.MyConf[Proto_SQL_serverdbrows] = config[Proto_Conf_serverdbrows4jfofflinefile]
		p.body.MyConf[Proto_SQL_clientsdbdriver] = config[Proto_Conf_clientsdbdriver4jfofflinefile]
		p.body.MyConf[Proto_SQL_clientsdbconn] = config[Proto_Conf_clientsdbconn4jfofflinefile]
		p.body.MyConf[Proto_SQL_clientsdbcol] = config[Proto_Conf_clientsdbcol4jfofflinefile]
		p.body.MyConf[Proto_SQL_clientsdbrows] = config[Proto_Conf_clientsdbrows4jfofflinefile]

	default:

		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Restart
		p.header.HeadMsg[Proto_Conf_clientstask] = Proto_Conf_Task_unknown
		return errors.New(Proto_Conf_Task_unknown)
	}
	return nil
}
