package main

import (
	"Zinx/src/zinx/zinterface"
	"Zinx/src/zinx/znet"
	"fmt"
)


//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

//处理conn业务之前的hook方法
func (this *PingRouter)PreHandle(request zinterface.IRequest){
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTcpConection().Write([]byte("befor ping\n"))
	if err != nil {
		fmt.Println("call back befor ping error")
	}
}

//处理conn业务的主方法
func (this *PingRouter)Handle(request zinterface.IRequest) {
	fmt.Println("Call Router Handle")
	_, err := request.GetConnection().GetTcpConection().Write([]byte("ping... ping\n"))
	if err != nil {
		fmt.Println("call back ping... ping error")
	}
}

//处理conn业务之后的子方法
func (this *PingRouter)PostHandle(request zinterface.IRequest){
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTcpConection().Write([]byte("after ping\n"))
	if err != nil {
		fmt.Println("call back after ping error")
	}
}


func main() {
	fmt.Println("start")
	s := znet.NewServer("[zinx v0.2]")

	//2 给当前zinx添加router
	s.AddRouter(&PingRouter{})
	//启动服务
	s.Server()
}
