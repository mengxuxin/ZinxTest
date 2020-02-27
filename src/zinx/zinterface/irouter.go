package zinterface

/*
路由抽象接口
路由里的数据都是IRequest
 */

type IRouter interface {
	//处理conn业务之前的hook方法
	PreHandle(request IRequest)

	//处理conn业务的主方法
	Handle(request IRequest)

	//处理conn业务之后的子方法
	PostHandle(request IRequest)
}
