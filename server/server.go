package server

import (
	"net"
	"time"
)

type MQServer struct {
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
