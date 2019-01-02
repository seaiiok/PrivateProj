package protocol

//Proto user protocol
type Protocol struct {
	Error error
	ProcessTrace []string
	MeConf map[string]string
	Data []string
	Datas [][]string
}