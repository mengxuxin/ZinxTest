package zinterface

type IServer interface {
	//start server
	Start()

	//stop server
	Stop()

	//go server
	Server()

	//路由功能，给当前的服务注册路由方法，供客户端的链接使用
	AddRouter(router IRouter)

}