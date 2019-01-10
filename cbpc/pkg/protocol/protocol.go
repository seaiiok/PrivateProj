package protocol

//Proto user protocol
type Proto struct {
	Err
	SQL
	Device
}

//Err user protocol
type Err struct {
	Error        bool
	ProcessTrace []string
}

//Device user protocol
type Device struct {
	DeviceName   string
	DeviceIP     string
	DeviceTask   string
	DeviceRouter string
}

//SQL user protocol
type SQL struct {
	DatabaseDriver string
	DatabaseSource string
	InsertSQL      string
	QuerySQL       string
	Args           []string
	Data           [][]string
}
