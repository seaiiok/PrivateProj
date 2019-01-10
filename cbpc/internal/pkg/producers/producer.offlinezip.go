package producers

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"ifix.cbpc/cbpc/pkg/conf"
	"ifix.cbpc/cbpc/pkg/db"
	"ifix.cbpc/cbpc/pkg/files"
	"ifix.cbpc/cbpc/pkg/protocol"
	"ifix.cbpc/cbpc/pkg/utils"
)

//OfflineZip ...
type OfflineZip struct {
	protocol.Proto
}
type filesdata struct {
	J_DeviceName   string
	J_DeviceIP     string
	J_FileName     string
	J_FilePath     string
	J_FileCreateD  string
	J_FileCreateDT string
	J_GoodSum      string
	J_BadSum       string
	J_StdNo        string
	J_PKNo         string
	J_Coder        string
	J_Dout         string
}

//SetObjectInfo ...
func (p *OfflineZip) SetObjectInfo() {
	p.SetErrorNil()
	p.Device.DeviceIP = utils.GetLocalMatchIp(conf.ConstMatchIP)
	p.Device.DeviceName = "检封离线压缩数据"
	p.Device.DeviceTask = conf.Config[conf.ConstClientsTask]
	p.Device.DeviceRouter = conf.ConstRouterGetObjectInfo
	p.SQL.DatabaseDriver = conf.Config[conf.ConstClientsDriver]
	p.SQL.DatabaseSource = conf.Config[conf.ConstClientsSource]
	fmt.Println(p)
}

//SetObjectConf ...
func (p *OfflineZip) SetObjectConf() {
	// fo := new(filesdata)
	// f := files.GetAllFiles()
	// path, err := f(p.SQL.DatabaseDriver)
	// res := files.GetNeedFilesInp(path, p.SQL.DatabaseSource)
	// res = files.GetNeedFilesMatch(res, conf.Config[conf.ConstClientsKeys])
	// fmt.Println(res[0])
	// for i := 0; i < len(res); i++ {
	// 	fileInfo, err := os.Stat(res[i])
	// 	if err != nil {
	// 		continue
	// 	}
	// 	fileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	// 	filecreatetimeN := fileSys.CreationTime.Nanoseconds()
	// 	tm := time.Unix(filecreatetimeN/1e9, 0)

	// 	fo.J_DeviceName = p.Device.DeviceName
	// 	fo.J_DeviceIP = p.Device.DeviceIP
	// 	fo.J_FileName = filepath.Base(res[i])
	// 	fo.J_FilePath = res[i]
	// 	fo.J_FileCreateD = tm.Format("2006-01-02")
	// 	fo.J_FileCreateDT = tm.Format("2006-01-02 15:04:05")
	// }

	p.SQL.DatabaseDriver = conf.Config[conf.ConstServerDriver]
	p.SQL.DatabaseSource = conf.Config[conf.ConstServerSource]
}

//SetObjectKeys ...
func (p *OfflineZip) SetObjectKeys() {
	var err error
	p.SQL.DatabaseDriver = conf.Config[conf.ConstClientsDriver]
	p.SQL.DatabaseSource = conf.Config[conf.ConstClientsSource]
	var fo []filesdata
	f := files.GetAllFiles()
	path, err := f(p.SQL.DatabaseDriver)
	res := files.GetNeedFilesInp(path, p.SQL.DatabaseSource)
	res = files.GetNeedFilesMatch(res, conf.Config[conf.ConstClientsKeys])
	fmt.Println("sum:",len(res))
	for i := 0; i < len(res); i++ {
		fo = make([]filesdata, len(res))
		fileInfo, err := os.Stat(res[i])
		if err != nil {
			continue
		}
		fileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
		filecreatetimeN := fileSys.CreationTime.Nanoseconds()
		tm := time.Unix(filecreatetimeN/1e9, 0)
		fo[i].J_FilePath = res[i]
		fo[i].J_FileCreateD = tm.Format("2006-01-02 15:04:05")
	}
	p.SQL.Args = make([]string, 0)
	for j := 0; j < len(fo); j++ {
		nowtime,_:=time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
		createtime,_:=time.Parse("2006-01-02 15:04:05", fo[j].J_FileCreateD)
		
		fmt.Println("createtime:",createtime)
		fmt.Println(nowtime.Sub(createtime).Hours())
		if nowtime.Sub(createtime).Hours()<=2160 {
			fmt.Println(nowtime.Sub(createtime).Hours())
		}
	}
	p.SQL.Args = make([]string, 0)
	p.SQL.Args = res
	if err != nil {
		p.SetError(true)
	}

	p.SQL.QuerySQL = conf.Config[conf.ConstClientsKeys]
	p.SQL.Args, err = db.Query(p.SQL.QuerySQL)
	if err != nil {
		p.SetError(true)
	}
	p.SQL.QuerySQL = conf.Config[conf.ConstServerKeys]
}

//SetObjectDatas ...
func (p *OfflineZip) SetObjectDatas() {
	var err error
	p.SQL.QuerySQL = conf.Config[conf.ConstClientsData]
	p.SQL.QuerySQL, err = getFulQuery(p.SQL.QuerySQL, p.SQL.Args)
	if err != nil {
		p.SetError(true)
	} else {
		p.SQL.Data, err = db.Querys(p.SQL.QuerySQL)
	}
	if err != nil {
		p.SetError(true)
	}
	p.SQL.InsertSQL = conf.Config[conf.ConstServerData]
}
