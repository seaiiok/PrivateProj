package net

import (
	"crypto/tls"
	"errors"
	"net/http"

	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/convert"
	"ifix.cbpc/cbpc/pkg/protocol"
)

//HTTPSURL client http url
var HTTPSURL = "https://" + conf.Config[conf.ConstServerAddr] + ":" + conf.Config[conf.ConstServerPort] + "/service"

//ClientsHTTPInit ...
func ClientsHTTPInit() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
}

//HTTPPost ...
func HTTPPost(p *protocol.Protocol) *protocol.Protocol {
	body, _ := convert.Struct2Reader(p)
	//create post req
	postReq, err := http.NewRequest(conf.ConstHTTPMethod,
		HTTPSURL,
		body,
	)
	if err != nil {
		p.Error = err
		return p
	}
	//add header
	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")
	//do post req
	resp, err := client.Do(postReq)
	if err != nil {
		p.Error = err
		return p
	}
	convert.Reader2Struct(resp.Body, p)
	if resp.StatusCode != 200 {
		p.Error = errors.New(resp.Status)
	}
	defer resp.Body.Close()
	return p
}
