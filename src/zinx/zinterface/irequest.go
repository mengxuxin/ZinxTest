package zinterface

/*
IREQuest 接口
把客户端请求的链接消息和请求的数据 包装到一个Request
 */

type IRequest interface {
	//得到当前链接
	GetConnection() IConnection

	//得到请求的消息数据
	GetData() []byte
}
