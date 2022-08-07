package server

import (
	"log"
	"net"
)

type IServer interface {
	Start()
	Stop()
}

type Server struct {
	*Config
	ConnectionManager *ConnectionManager // 链接集中管理
}

func NewServer(config Config) IServer {
	s := &Server{
		Config: &Config{
			MaxConnectionCount: config.MaxConnectionCount,
			Network:            config.Network,
			Address:            config.Address,
			Port:               config.Port,
		},
	}

	s.ConnectionManager = NewConnectionManager(s)

	return s
}

func (s *Server) Start() {
	listener, err := net.Listen(s.Config.Network, s.Config.Address+":"+s.Config.Port)
	if err != nil {
		panic(err)
	}
	log.Printf("服务器启动成功，正在监听 %s:%s\n", s.Config.Address, s.Port)
	for {
		conn, err := listener.Accept()
		if err != nil {

			continue
		}

		// 对于一个客户端的链接，这是开启的第一个协程
		go s.ConnectionManager.HandleConnection(conn)
	}
}

func (s *Server) Stop() {

}
