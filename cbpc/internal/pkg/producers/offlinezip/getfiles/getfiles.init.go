package getfiles


type GetFiles interface{
	GetAllFiles(string,string,string)
	GetFilesInfo(string,string)
	GetFilesData([]string,string,string)
}