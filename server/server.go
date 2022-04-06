package server

import (
	"net"
	"time"
)

type MQServer struct {
	// 客户端ID递增序列
	ClientIDSequence int64
	// 服务端启动时间
	startTime time.Time
	tcpServer *tcpServer

	tcpListener net.Listener
}

func New() (*MQServer, error) {
	var err error
	s := &MQServer{
		startTime: time.Now(),
	}

	s.tcpServer = &tcpServer{server: s}
	s.tcpListener, err = net.Listen("tcp", "0.0.0.0:6789")

	return s, err
}

func (s *MQServer) Main() error {
	var err error
	err = TCPServer(s.tcpListener, s.tcpServer)

	return err
}
