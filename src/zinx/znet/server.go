package znet

import (
	"Zinx/src/zinx/utils"
	"Zinx/src/zinx/zinterface"
	"errors"
	"fmt"
	"net"
)

type Server struct {
	//name
	Name string

	//ip version
	IPVersion string

	//ip
	IP string

	//port
	nPort int

	//当前的server添加一个router
	Router zinterface.IRouter
}

//定义当前客户端链接所绑定的handle.api(目前这个handle是写死的，以后优化)
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int)error {
	//处理回显
	fmt.Println("[conn Handle] CallBackToClient")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("callbacktoclient error")
	}

	return nil
}

//start
func (s *Server) Start() {
	fmt.Printf("[zinx]server name: %s, listerner at ip:%s, port: %d is starting",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[zinx] Version: %s, MaxConn: %d, MaxPackageSize:%d\n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	fmt.Printf("start server listener at IP:%s, port:%d", s.IP, s.nPort)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.nPort))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			panic(err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen :", s.IPVersion, " err: ", err)
		}
		fmt.Println("start zinx server success: ",s.Name, "listening...")
		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			//v0.1
			//go func() {
			//	for {
			//		buf := make([]byte, 512)
			//		cnt, err := conn.Read(buf)
			//		if err != nil {
			//			fmt.Println("recv buf err", err)
			//			continue
			//		}
			//		fmt.Printf("server recvie %s\n", buf[:cnt])
			//		if _,err := conn.Write(buf[:cnt]); err != nil {
			//			fmt.Println("write back buf err", err)
			//			continue
			//		}
			//	}
			//}()
			//v0.2
			//将处理新连接的业务方法和conn 进行绑定 得到我们的链接模块
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			go dealConn.Start()
		}
	}()
}

//stop
func (s *Server) Stop() {

}

//server
func (s *Server) Server() {
	s.Start()

	//需要阻塞住
	select {

	}
}

func (s *Server) AddRouter(router zinterface.IRouter) {
	s.Router = router
	fmt.Println("add router success")
}

//init server
func NewServer(name string) zinterface.IServer {
	s := &Server {
		Name :utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:utils.GlobalObject.Host,
		nPort:utils.GlobalObject.TcpPort,
		Router:nil,
	}


	return s
}