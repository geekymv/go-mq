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

	topicMap map[string]*Topic
}

func New() (*MQServer, error) {
	var err error
	s := &MQServer{
		startTime: time.Now(),
		topicMap:  make(map[string]*Topic),
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

func (s *MQServer) GetTopic(topicName string) *Topic {

	t, ok := s.topicMap[topicName]
	if ok {
		return t
	}

	t = NewTopic(topicName, s)
	s.topicMap[topicName] = t

	t.Start()

	return t
}
