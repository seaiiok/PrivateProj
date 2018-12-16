package api

type AppInfo struct {
	ServerInfo  Info
	ClientsInfo Info
}

type Info struct {
	Name      string
	Count     int
	RunInfo   []string
	StartTime string
	Doing     []string
}

func (app *AppInfo) ApiServerInfo() Info {
	return app.ServerInfo
}

func (app *AppInfo) ApiClientsInfo() Info {
	return app.ClientsInfo
}
