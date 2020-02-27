package znet

import (
	"Zinx/src/zinx/zinterface"
	"fmt"
	"net"
)

type Connection struct {
	//current socket TCP
	Conn *net.TCPConn

	//connection id
	ConnID uint32

	//status
	isClosed bool

	//v0.2
	//当前链接绑定的处理业务方法API
	//handleAPI zinterface.HandleFunc

	//告知当前连接已经停止的channel
	ExitChan chan bool

	//v0.3
	//该链接处理的方法Router
	Router zinterface.IRouter
}

func (c *Connection)StartReader()	{
	fmt.Println("Reader Goroutine Start")
	fmt.Println("connID = ", c.ConnID, "Reader is exit, remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err",err)
			continue
		}

		//v0.2
		////调用当前链接所绑定的api， 回调如何处理交由回调函数去整
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("connID", c.ConnID, "handle is error",err)
		//	break
		//}

		//v0.3
		//得到当前conn数据的Request请求数据
		req := Request{
			conn:c,
			data:buf,
			}

		//执行注册的路由方法
		go func(request zinterface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
		//从路由中找到对应的conn对应的router
	}
}

func (c *Connection) Start() {
	fmt.Println("conn start()... connID=", c.ConnID)

	//启动读数据的业务
	go c.StartReader()

	//启动写数据的业务
}

func (c *Connection)Stop()	{
	fmt.Println("Conn Stop()...connID=", c.ConnID)

	if c.isClosed == true {
		return
	}

	//回收资源
	c.Conn.Close()

	close(c.ExitChan)

	//状态
	c.isClosed = true
}

//get socket conn
func (c *Connection)GetTcpConection() *net.TCPConn {
	return c.Conn
}

//get socker connID
func (c *Connection)GetConnID()	uint32 {
	return c.ConnID
}

//get client tcp status,IP port
func (c *Connection)RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//send data to remote client
func (c *Connection)Send(data []byte) error {
	return nil
}

//初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router zinterface.IRouter) zinterface.IConnection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		//handleAPI: callback_api,
		Router:router,
		ExitChan:  make(chan bool, 1),
	}
	return c
}