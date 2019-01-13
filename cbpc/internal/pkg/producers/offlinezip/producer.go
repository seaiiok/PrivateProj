package producer

import (
	"ifix.cbpc/cbpc/pkg/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"ifix.cbpc/cbpc/internal/pkg/producers/offlinezip/getfiles"
	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/convert"
	"ifix.cbpc/cbpc/pkg/protocol"
)

//OffilneZip ...
type OffilneZip struct {
	protocol.Proto
}

//SetObjectInfo ...
func (p *OffilneZip) SetObjectInfo() {
	days, err := strconv.Atoi(conf.Config[conf.ConstClientsData])
	if err != nil {
		p.Err.Error = true
	}
	p.Device.DeviceIP = utils.GetLocalMatchIp(conf.ConstMatchIP)
	p.Device.DeviceName = "检封工房离线数据压缩包"
	p.Device.DeviceRouter = conf.ConstRouterGetObjectInfo
	p.Device.DeviceTask = conf.Config[conf.ConstClientsTaskOfflineZip]
	p.SQL.DatabaseDriver=conf.Config[conf.ConstClientsDriver]
	p.SQL.DatabaseSource=conf.Config[conf.ConstClientsSource]
	p.SQL.QuerySQL = conf.Config[conf.ConstClientsKeys]
	p.SQL.InsertSQL = strconv.Itoa(days)
}

//SetObjectPost ...
func (p *OffilneZip) SetObjectPost(client *http.Client) {

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
func (p *OffilneZip) SetObjectConf() {
	// days, _ := strconv.Atoi(p.SQL.InsertSQL)
	// fo := getLastNeedFilesInfo(p.SQL.DatabaseDriver, p.SQL.DatabaseSource, p.SQL.QuerySQL, p.Device.DeviceName,p.Device.DeviceIP , days)
	
}

//SetObjectKeys ...
func (p *OffilneZip) SetObjectKeys() {
	p.SQL.Args=make([]string,0)
	days, _ := strconv.Atoi(p.SQL.InsertSQL)
	fo := getLastNeedFilesInfo(p.SQL.DatabaseDriver, p.SQL.DatabaseSource, p.SQL.QuerySQL, p.Device.DeviceName,p.Device.DeviceIP , days)
	for i := 0; i < len(fo.FilesInfo); i++ {
		p.SQL.Args=append(p.SQL.Args,fo.FilesInfo[i].J_FilePath)
	}
	
}

//SetObjectDatas ...
func (p *OffilneZip) SetObjectDatas() {

}

//getLastNeedFilesInfo ...
func getLastNeedFilesInfo(rootpath, inp, matexfile, name, matip string, days int) getfiles.FilesInfo {
	var gf getfiles.GetFiles
	fs := new(getfiles.FilesInfo)
	res := new(getfiles.FilesInfo)
	gf = fs
	gf.GetAllFiles(rootpath, inp, matexfile)
	gf.GetFilesInfo(name, matip)
	lastdt := time.Now()
	dt := lastdt.Add(time.Duration(days*24*-1) * time.Hour)
	for i := 0; i < len(fs.FilesInfo); i++ {
		if fs.FilesInfo[i].J_FileCreateD >= dt.Format("2006-01-02") {
			res.FilesInfo = append(res.FilesInfo, fs.FilesInfo[i])
		}
	}
	return *res
}

//getLastFilesInfo ...
func getNeedFilesInfo(rootpath, inp, matexfile, name, matip string) getfiles.FilesInfo {
	var gf getfiles.GetFiles
	fs := new(getfiles.FilesInfo)
	gf = fs
	gf.GetAllFiles(rootpath, inp, matexfile)
	gf.GetFilesInfo(name, matip)
	return *fs
}

//getFilesData ...
func getFilesData(paths []string, ip, name string) [][]string {
	var gf getfiles.GetFiles
	res := make([][]string, 0)
	fs := new(getfiles.FilesInfo)
	gf = fs
	gf.GetFilesData(paths, ip, name)
	for i := 0; i < len(fs.FilesInfo); i++ {
		re := make([]string, 0)

		// fs.FilesInfo[i].FileData
		rd := getfiles.GetRowsData(fs.FilesInfo[i].FileData)
		for j := 0; j < len(rd); j++ {
			re = append(re, fs.FilesInfo[i].J_DeviceName, fs.FilesInfo[i].J_DeviceIP, fs.FilesInfo[i].J_FileName, fs.FilesInfo[i].J_FilePath, fs.FilesInfo[i].J_FileCreateD, fs.FilesInfo[i].J_FileCreateDT)
			re = append(rd[j])
			res = append(res, re)
		}
	}
	return res
}
