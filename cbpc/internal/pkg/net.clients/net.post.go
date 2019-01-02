package net

import (
	"crypto/tls"
	"errors"
	"net/http"

	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/convert"
)

//HTTPSURL client http url
var HTTPSURL = "https://" + conf.Config["serveraddr"] + ":" + conf.Config["serverport"] + "/service"

//ClientsHTTPInit ...
func ClientsHTTPInit() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
}

//HTTPPost ...
func (p *ClientsProto) HTTPPost() {
	if p.Error != nil {
		p.SetProcessTrace()
	}

	body, _ := convert.Struct2Reader(p)
	//create post req
	postReq, err := http.NewRequest(conf.ConstHTTPMethod,
		HTTPSURL,
		body,
	)
	if err != nil {
		p.SetProcessTrace(err.Error())
		p.Error = err
		return
	}
	//add header
	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")
	//do post req
	resp, err := client.Do(postReq)
	if err != nil {
		p.SetProcessTrace(err.Error())
		p.Error = err
		return
	}
	convert.Reader2Struct(resp.Body, p)
	if resp.StatusCode != 200 {
		p.Error = errors.New(resp.Status)
	}
	defer resp.Body.Close()
}
