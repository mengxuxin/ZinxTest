package znet

import "Zinx/src/zinx/zinterface"

type Request struct {
	//已经和客户端建立好链接的conn
	conn zinterface.IConnection

	//客户端请求的data
	data []byte
}

func (r *Request) GetConnection() zinterface.IConnection {
	return r.conn
}


func (r *Request) GetData() []byte {
	return r.data
}