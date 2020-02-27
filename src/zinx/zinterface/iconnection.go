package zinterface

import "net"

type IConnection interface {
	//start connection
	Start()

	//stop connection
	Stop()

	//get socket conn
	GetTcpConection() *net.TCPConn

	//get socker connID
	GetConnID()	uint32

	//get client tcp status,IP port
	RemoteAddr() net.Addr

	//send data to remote client
	Send(data []byte) error
}

//处理链接业务的方法
type HandleFunc	func(*net.TCPConn, []byte, int) error
