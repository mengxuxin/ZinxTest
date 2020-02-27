package utils

import (
	"Zinx/src/zinx/zinterface"
	"encoding/json"
	"io/ioutil"
)

/*
存储一切有关Zinx框架的全局参数，供其他模块使用
	参数可通过Zinx.json 由用户进行配置
 */

type GlobalObj struct {
	//Server
	TcpServer zinterface.IServer //当前Zinx 全局的Server对象
	Host string					 //服务器主机监听的ip
	TcpPort int					 //
	Name string					//服务器的名称

	//Zinx
	Version        string //当前的版本号
	MaxConn        int    //当前的服务器主机允许的最大连接数
	MaxPackageSize uint32 //当前Zinx框架数据包的最大值

}

var GlobalObject *GlobalObj

/*
提供初始化方法，给定初始值
 */
func init()  {
	//如果没有加载配置文件，执行默认值
	GlobalObject = &GlobalObj{
		TcpServer:      nil,
		Host:           "0.0.0.0",
		TcpPort:        0,
		Name:           "ZinxServerApp",
		Version:        "V0.4",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	//应该尝试从conf/zinx.json 加载用户自定义的参数
	GlobalObject.ReloadFromJson()
}


/*
从 zinx.json 加载自定义的参数
 */
func (g *GlobalObj)ReloadFromJson() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	//将json 文件解析到结构体中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}