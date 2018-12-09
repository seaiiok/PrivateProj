package service

type ServerInterface interface {
	GetDBReady()
	GetDBKeyArray()
	SetDBData()
}