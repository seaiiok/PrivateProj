/*
Package producers its a collector for cbpm. 

All data collector should be implements interface "Producers" or  "Object". see example!

here collectors for "成都印钞厂MES数据采集项目"

user protocol error should use SetError method ,or headache!
oops! :)

example:


//Example ...
type Example struct {
	protocol.Proto
}

//SetObjectInfo ...
func (p *Example) SetObjectInfo() {

}

//SetObjectPost ...
func (p *Example) SetObjectPost(client *http.Client) {

	URL := "https://" + conf.Config[conf.ConstServerAddr] + ":" + conf.Config[conf.ConstServerPort] + p.Device.DeviceRouter

	method := "POST"

	body, _ := convert.Struct2Reader(p)

	//create post req
	postReq, err := http.NewRequest(
		method,
		URL,
		body,
	)
	if err != nil {
		p.SetError(true)
		return
	}
	//add header
	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")
	//do post req
	resp, err := client.Do(postReq)
	if err != nil {
		p.SetError(true)
		p.SetProcessTrace(err.Error())
		return
	}
	p.DeviceStatusCode = resp.StatusCode
	convert.Reader2Struct(resp.Body, p)
	defer resp.Body.Close()
}

//SetObjectConf ...
func (p *Example) SetObjectConf() {

}

//SetObjectKeys ...
func (p *Example) SetObjectKeys() {

}

//SetObjectDatas ...
func (p *Example) SetObjectDatas() {

}


*/
package producers