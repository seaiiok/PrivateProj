package getfiles

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"ifix.cbpc/cbpc/pkg/files"
)

type FileInfo struct {
	J_DeviceName   string
	J_DeviceIP     string
	J_FileName     string
	J_FilePath     string
	J_FileCreateD  string
	J_FileCreateDT string
	FileData       []string
}

type FilesInfo struct {
	FilesInfo []FileInfo
}

func (f *FilesInfo) GetAllFiles(rootpath string, inp string, mat string) {
	f.FilesInfo = make([]FileInfo, 0)
	fi := files.GetAllFiles()
	paths, err := fi(rootpath)
	if err != nil {
		return
	}
	pathsn := files.GetNeedFilesInp(paths, inp)
	pathsnew := files.GetNeedFilesMatch(pathsn, mat)
	for i := 0; i < len(pathsnew); i++ {
		fileinfo := new(FileInfo)
		fileinfo.J_FilePath = pathsnew[i]
		f.FilesInfo = append(f.FilesInfo, *fileinfo)
	}
}

func (f *FilesInfo) GetFilesInfo(name string, ip string) {
	for i := 0; i < len(f.FilesInfo); i++ {
		// f.FilesInfo[i].J_FilePath

		fs, err := os.Stat(f.FilesInfo[i].J_FilePath)
		if err != nil {
			continue
		}

		fileSys := fs.Sys().(*syscall.Win32FileAttributeData)
		filecreatetimeN := fileSys.CreationTime.Nanoseconds()
		tm := time.Unix(filecreatetimeN/1e9, 0)

		f.FilesInfo[i].J_DeviceIP = ip
		f.FilesInfo[i].J_DeviceName = name
		f.FilesInfo[i].J_FileCreateD = tm.Format("2006-01-02")
		f.FilesInfo[i].J_FileCreateDT = tm.Format("2006-01-02 15:04:05")
		f.FilesInfo[i].J_FileName = filepath.Base(f.FilesInfo[i].J_FilePath)

	}
}

func (f *FilesInfo) GetFilesData(paths []string, ip string, name string) {
	f.FilesInfo = make([]FileInfo, 0)
	for i := 0; i < len(paths); i++ {

		fs, err := os.Stat(paths[i])
		if err != nil {
			continue
		}
		fileSys := fs.Sys().(*syscall.Win32FileAttributeData)
		filecreatetimeN := fileSys.CreationTime.Nanoseconds()
		tm := time.Unix(filecreatetimeN/1e9, 0)

		f.FilesInfo[i].J_DeviceIP = ip
		f.FilesInfo[i].J_DeviceName = name
		f.FilesInfo[i].J_FileCreateD = tm.Format("2006-01-02")
		f.FilesInfo[i].J_FileCreateDT = tm.Format("2006-01-02 15:04:05")
		f.FilesInfo[i].J_FileName = filepath.Base(paths[i])

		f.FilesInfo[i].FileData = make([]string, 0)
		f.FilesInfo[i].FileData = files.ReadZipLines(paths[i])
	}
}

func GetRowsData(row []string) [][]string {
	res := make([][]string, 0)
	GoodSum := 0
	BadSum := 0

	for i := 0; i < len(row); i++ {
		re := make([]string, 0)
		rowarr := strings.Split(row[i], " ")
		lenrow := len(rowarr)
		if lenrow >= 4 {
			StdNo := rowarr[0]
			PKNo := rowarr[1]
			Coder := rowarr[lenrow-2]
			Dout := rowarr[lenrow-1]
			if len(StdNo) >= 12 && len(PKNo) >= 10 && len(Coder) == 2 && len(Dout) == 2 {
				if Coder == "80" {
					GoodSum++
				} else {
					BadSum++
				}
				re = append(re, strconv.Itoa(GoodSum), strconv.Itoa(BadSum), StdNo, PKNo, Coder, Dout)
			}
		}
		if len(re) == 6 {
			res = append(res, re)
		}
	}
	return res
}
