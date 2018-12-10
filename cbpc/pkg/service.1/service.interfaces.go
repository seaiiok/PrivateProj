package service

// IServerDatebase 服务器数据库接口，本地mssql2012
type IServerDatebase interface{
	ServerDatebaseGetReady(req Proto) (res Proto)
	ServerDatebaseGetData(req Proto) (res Proto)
	ServerDatebaseSetData(req Proto) (res Proto)
}

// IClientsDatebase 客服端数据库接口，3#门车辆磅秤，离线检测
type IClientsDatebase interface {
	ClientsDatebaseGetReady(req Proto) (res Proto)
	ClientsDatebaseGetData(req Proto) (res Proto)
}

// IClientsOpc 客服端OPC接口，集中空调，集中供液等
type IClientsOpc interface {
	ClientsDatebaseGetReady(req Proto) (res Proto)
	ClientsDatebaseGetData(req Proto) (res Proto)
}

// IClientsFiles 客服端文件接口,在线清数日志，检封工坊离线数据包
type IClientsFiles interface {
	ClientsFilesGetReady(req Proto) (res Proto)
	ClientsFilesGetData(req Proto) (res Proto)
}

// IServer 服务端超级接口
type IServer interface {
	IServerDatebase
}

// IClients 客服端超级接口
type IClients interface {
	IClientsDatebase
	IClientsFiles
}

// IService 项目接口
type IService interface {
	IServer
	IClients
}