/*
Package producers its a collector for cbpm. 

All data collector should be implements interface "Producers" or  "Object". see example!

here collectors for "成都印钞厂MES数据采集项目"

user protocol error should use SetError method ,or headache!
oops! :)

example:

//Objects Producers ...
type Objects struct {
	import "ifix.cbpc/cbpc/pkg/protocol"
	use this user protocol
}

//SetObjectInfo ...
func (p *Objects) SetObjectInfo() {
//here set device info
}

//SetObjectConf ...
func (p *Objects) SetObjectConf() {
//here set device config
}

//SetObjectKeys ...
func (p *Objects) SetObjectKeys() {
//here set device data keys
}

//SetObjectDatas ...
func (p *Objects) SetObjectDatas() {
//here set device data
}

*/
package producers