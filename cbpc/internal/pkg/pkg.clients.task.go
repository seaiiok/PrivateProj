package pkg

import (
	"errors"
)

var iProducte IProducte

//GetClientsConf config for all clients
func (p *Proto) GetClientsConf() error {

	switch p.header.HeadMsg[Proto_Conf_clientstask] {
	case Proto_Conf_Task_d3weight:
		c := new(clientsd3weight)
		c.clients.MyConf = p.body.MyConf
		for k, v := range p.body.MyConf {
			config[k] = v
		}
		iProducte = c
		err := iProducte.ClientsProducterGetReady()
		if err != nil {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Restart
		}
	case Proto_Conf_Task_jwofflinedb:

	case Proto_Conf_Task_jwonlinefile:

	case Proto_Conf_Task_jfofflinedb:

	case Proto_Conf_Task_jfofflinefile:
		c := new(clientsd3weight)
		c.clients.MyConf = p.body.MyConf
		for k, v := range p.body.MyConf {
			config[k] = v
		}
		iProducte = c
		err := iProducte.ClientsProducterGetReady()
		if err != nil {
			p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Restart
		}
	default:

		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Restart
		p.header.HeadMsg[Proto_Conf_clientstask] = Proto_Conf_Task_unknown
		return errors.New(Proto_Conf_Task_unknown)
	}
	return nil
}

//GetClientsCol col for all clients
func (p *Proto) GetClientsCol() error {

	switch p.header.HeadMsg[Proto_Conf_clientstask] {
	case Proto_Conf_Task_d3weight:
		p.sqlStringMakeClientsDTimes(config[Proto_SQL_clientsdbcol])
		c := new(clientsd3weight)
		c.clients.MyConf = make(map[string]string)
		c.clients.MyConf = p.body.MyConf
		iProducte = c
		p.body.BodyData = make([]string, 0)
		res := iProducte.ClientsProducterGetData()
		if len(res) > 0 {
			for i := 0; i < len(res); i++ {
				p.body.BodyData = append(p.body.BodyData, res[i])
			}

		}

	case Proto_Conf_Task_jwofflinedb:

	case Proto_Conf_Task_jwonlinefile:

	case Proto_Conf_Task_jfofflinedb:

	case Proto_Conf_Task_jfofflinefile:

	default:

		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Restart
		p.header.HeadMsg[Proto_Conf_clientstask] = Proto_Conf_Task_unknown
		return errors.New(Proto_Conf_Task_unknown)
	}
	return nil
}

//GetClientsRows Rows for all clients
func (p *Proto) GetClientsRows() error {

	switch p.header.HeadMsg[Proto_Conf_clientstask] {
	case Proto_Conf_Task_d3weight:

		p.sqlStringClientsMakeRows(config[Proto_SQL_clientsdbrows], p.body.BodyData)

		c := new(clientsd3weight)
		c.clients.MyConf = p.body.MyConf

		iProducte = c

		res := iProducte.ClientsProducterGetDatas()
		p.body.BodyDatas = make([][]string, 0)
		if len(res) > 0 {
			for i := 0; i < len(res); i++ {
				p.body.BodyDatas = append(p.body.BodyDatas, res[i])
			}
		}

	case Proto_Conf_Task_jwofflinedb:

	case Proto_Conf_Task_jwonlinefile:

	case Proto_Conf_Task_jfofflinedb:

	case Proto_Conf_Task_jfofflinefile:

	default:

		p.header.HeadMsg[Proto_Cmd] = Proto_Cmd_Restart
		p.header.HeadMsg[Proto_Conf_clientstask] = Proto_Conf_Task_unknown
		return errors.New(Proto_Conf_Task_unknown)
	}
	return nil
}
