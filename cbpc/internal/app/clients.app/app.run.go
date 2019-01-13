package app

import (
	"fmt"
	"reflect"
	"time" // onlinelog "ifix.cbpc/cbpc/internal/pkg/producers/onlinelog"

	d3weight "ifix.cbpc/cbpc/internal/pkg/producers/d3weight"
	offlinezip "ifix.cbpc/cbpc/internal/pkg/producers/offlinezip"
	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/interfaces"
)

func clientsServiceStart() {
	go clientsStart()
}

// ServiceStart ...
func clientsStart() {
	var o interfaces.Producers
	reconn := 3
	switch conf.Config[conf.ConstClientsTask] {
	case conf.ConstClientsTaskD3Weight:
		o = new(d3weight.D3weight)
	case conf.ConstClientsTaskOnlineLog:
		// o = new(onlinelog.OnlineLog)
	case conf.ConstClientsTaskOfflineZip:
		o = new(offlinezip.OffilneZip)
	default:
	}

	o.SetObjectInfo()
	for {
		rv := reflect.ValueOf(o).Elem()
		dr := rv.FieldByName(conf.ConstHTTPRouter).String()
		deviceStatusCode := rv.FieldByName(conf.ConstHTTPStatusCode).Int()
		for deviceStatusCode == conf.ConstHTTPStatusCodeOK {
			fmt.Println(deviceStatusCode)
			fmt.Println(dr)
			switch dr {

			case conf.ConstRouterGetObjectInfo:
				o.SetObjectInfo()
			case conf.ConstRouterGetObjectConf:
				o.SetObjectConf()
			case conf.ConstRouterGetObjectKeys:
				o.SetObjectKeys()
			case conf.ConstRouterGetObjectDatas:
				o.SetObjectDatas()
			case conf.ConstRouterGetObjectAgain:
				// time.Sleep(conf.ConstWaitTime * time.Second)
				o.SetObjectConf()
			case conf.ConstRouterGetObjectRestart:
				time.Sleep(conf.ConstWaitTime * time.Second)
				o.SetObjectInfo()
			default:
				time.Sleep(conf.ConstWaitTime * time.Second)
				o.SetObjectInfo()
			}

			o.SetObjectPost(client)
			dr = rv.FieldByName(conf.ConstHTTPRouter).String()
			deviceStatusCode = rv.FieldByName(conf.ConstHTTPStatusCode).Int()
		}

		if deviceStatusCode != conf.ConstHTTPStatusCodeOK {
			time.Sleep(time.Duration(reconn*conf.ConstWaitTime) * time.Second)
			reconn++
			if reconn >= 3 {
				rv.FieldByName(conf.ConstHTTPStatusCode).SetInt(conf.ConstHTTPStatusCodeOK)
				reconn = 0
			}
		}
	}

}
