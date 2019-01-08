package net

import (
	"crypto/tls"
	"net/http"
	"reflect"

	"ifix.cbpc/cbpc/interfaces"
	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/convert"
)

var client *http.Client

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
}

//NetPost ...
func NetPost(p *interfaces.Object) int {
	body, _ := convert.Struct2Reader(p)
	r := reflect.ValueOf(*p).Elem()
	v := r.FieldByName("DeviceRouter").String()
	URL := "https://" + conf.Config[conf.ConstServerAddr] + ":" + conf.Config[conf.ConstServerPort] + "/" + v
	method := "POST"
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
