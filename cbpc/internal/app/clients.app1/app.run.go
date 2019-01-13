package app

import (
	"fmt"
	"reflect"
	"time"

	d3weight "ifix.cbpc/cbpc/internal/pkg/producers/d3weight"
	offlinezip "ifix.cbpc/cbpc/internal/pkg/producers/offlinezip"
	onlinelog "ifix.cbpc/cbpc/internal/pkg/producers/onlinelog"
	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/interfaces"
)

func clientsServiceStart() {
	var statuscode int = 200
	go clientsStart(&statuscode)
	go reconn(&statuscode)

}

// ServiceStart ...
func clientsStart(statuscode *int) {
	var o interfaces.Producers

	switch conf.Config[conf.ConstClientsTask] {
	case conf.ConstClientsTaskD3Weight:
		o = new(d3weight.D3weight)
	case conf.ConstClientsTaskOnlineLog:
		o = new(onlinelog.OnlineLog)
	case conf.ConstClientsTaskOfflineZip:
		o = new(offlinezip.OfflineZip)
	default:
	}

	o.SetObjectInfo()
	for {
		for *statuscode == 200 {

			prouter := reflect.ValueOf(o).Elem()
			proute := prouter.FieldByName("DeviceRouter").String()
			pro := prouter.FieldByName("ProcessTrace")
			fmt.Println(pro)
			switch proute {
			case conf.ConstRouterGetObjectInfo:
				o.SetObjectInfo()
			case conf.ConstRouterGetObjectConf:
				o.SetObjectConf()
			case conf.ConstRouterGetObjectKeys:
				o.SetObjectKeys()
			case conf.ConstRouterGetObjectDatas:
				o.SetObjectDatas()
			case conf.ConstRouterGetObjectAgain:
				time.Sleep(conf.ConstWaitTime * time.Second)
				o.SetObjectConf()
			case conf.ConstRouterGetObjectRestart:
				time.Sleep(conf.ConstWaitTime * time.Second)
				o.SetObjectInfo()

			default:
				fmt.Println("default")
			}

			*statuscode = NetPost(&o)
		}
	}

}

func reconn(statuscode *int) {
	for {
		// statuscode.reConnSum++

		if *statuscode != 200 {
			time.Sleep(conf.ConstWaitTime * time.Second)
			*statuscode = 200
			// statuscode.reConnSum = 0
		}
		time.Sleep(10 * conf.ConstWaitTime * time.Second)
	}
}
