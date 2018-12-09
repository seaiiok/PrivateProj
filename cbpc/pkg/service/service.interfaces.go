package service

type ServerInterface interface {
	//服务就绪
	ServerServiceReady(req MyProto) (res MyProto)
	//响应资源
	ServerServiceRes(req MyProto) (res MyProto)
	//执行请求
	ServerServiceExec(req MyProto) (res MyProto)
}

type ClientsInterface interface {
	//服务就绪
	ClientsServiceReady(req MyProto) (res MyProto)
	//响应资源
	ClientsServiceRes(req MyProto) (res MyProto)
	//执行请求
	ClientsServiceExec(req MyProto) (res MyProto)
}

type ServiceInterface interface {
	ServerService ServerInterface
	ClientsService ClientsInterface
}