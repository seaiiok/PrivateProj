package main

import (

	"fmt"

	"ifix.cbpc/cbpc/api/server.api"
	 "ifix.cbpc/cbpc/pkg/conf"
)

func main() {


	cli1 := new(api.D3Weight)
	cli1.DriverName = conf.Config[conf.ConstServerDriverName]
	cli1.DataSourceName = conf.Config[conf.ConstServerDataSourceName]
	var c1 api.Consumer = cli1
	fmt.Println(cli1)
	err := api.GetConsumerPing(c1)
	fmt.Println(err)
	
	cli2 := cli1
	cli2.SqlSelectCols = conf.Config[conf.ConstServerD3WeightCols]
	var c2 api.Consumer = cli2
	fmt.Println(cli2.SqlSelectCols)
	err = api.GetConsumerKeys(c2)
	fmt.Println(err)
	fmt.Println(c2)

}
