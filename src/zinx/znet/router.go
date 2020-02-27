package znet

import "Zinx/src/zinx/zinterface"

/*
消息 对应路由方式
路由的接口
 */

type BaseRouter struct {

}

//多态的方式。先实现基类，有子类需要的时候就可以单独实现具体的方法

func (br *BaseRouter) PreHandle(request zinterface.IRequest) {}

func (br *BaseRouter) Handle(request zinterface.IRequest) {}

func (br *BaseRouter) PostHandle(request zinterface.IRequest) {}