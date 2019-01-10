package app

import (
	"crypto/tls"
	"net/http"
	"reflect"

	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/convert"
	"ifix.cbpc/cbpc/pkg/interfaces"
)

var client *http.Client

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
}

//NetPost ...
func NetPost(p *interfaces.Producers) int {
	pr := reflect.ValueOf(*p).Elem()
	// perr := pr.FieldByName("Error").IsNil()
	// if perr != true {
	// 	pr.FieldByName("DeviceRouter").SetString(conf.ConstRouterGetObjectRestart)
	// }

	URL := "https://" + conf.Config[conf.ConstServerAddr] + ":" + conf.Config[conf.ConstServerPort] + pr.FieldByName("DeviceRouter").String()

	method := "POST"

	body, _ := convert.Struct2Reader(*p)

	//create post req
	postReq, err := http.NewRequest(
		method,
		URL,
		body,
	)
	if err != nil {
		return 404
	}
	//add header
	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")
	//do post req
	resp, err := client.Do(postReq)
	if err != nil {
		return 404
	}

	convert.Reader2Struct(resp.Body, p)
	defer resp.Body.Close()
	return resp.StatusCode
}
