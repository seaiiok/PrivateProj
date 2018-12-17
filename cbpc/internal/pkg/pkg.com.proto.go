package pkg

type header struct {
	//Err...
	Error error
	//请求资源
	HeadMsg map[string]string
	//进程追溯
	ProcessTrace []string
}

type body struct {
	MyConf map[string]string
	BodyData []string
	BodyDatas [][]string
}

//Proto 协议
type Proto struct {
	header
	body
}
