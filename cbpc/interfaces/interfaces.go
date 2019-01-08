package interfaces

// Object collector interface
type Object interface {
	Consumers
	Producers
}

// Consumers server interface
type Consumers interface {
	GetObjectInfo()
	GetObjectConf()
	GetObjectKeys()
	GetObjectDatas()
}

// Producers clients interface
type Producers interface {
	SetObjectInfo()
	SetObjectConf()
	SetObjectKeys()
	SetObjectDatas()
}
