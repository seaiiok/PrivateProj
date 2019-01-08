package proto

//Proto user protocol
type Proto struct {
	Err
	Config
	Sql
	Device
}

//Err user protocol
type Err struct {
	Err          error
	processTrace []string
}

//config user protocol
type Config struct {
	Config map[string]string
}

//Device user protocol
type Device struct {
	DeviceName   string
	DeviceIP     string
	DeviceTask   string
	DeviceRouter string
}

//Sql user protocol
type Sql struct {
	DatabaseDriver string
	DatabaseSource string
	Insert         string
	Query          string
	Args           []string
	Data           [][]string
}
