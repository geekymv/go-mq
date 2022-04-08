package server

import (
	"io"
	"log"
	"net"
	"sync"

	"github.com/geekymv/go-mq/internal/protocol"
)

type tcpServer struct {
	server *MQServer
	// 存储连接
	conns sync.Map
}

func (s *tcpServer) Handle(conn net.Conn) {
	log.Printf("remote %v\n", conn.RemoteAddr())

	// 读取协议版本号
	buf := make([]byte, 2)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		log.Printf("failed to read protocol version -%v", err)
		conn.Close()
		return
	}
	v := string(buf)
	log.Printf("client version:%v\n", v)
	var prot protocol.Protocol
	switch v {
	case protocol.MagicV1:
		prot = &protocolV1{server: s.server}
	default:
		conn.Close()
		log.Printf("client(%s) error version '%s'\n", conn.RemoteAddr(), v)
		return
	}
	// 创建 client
	client := prot.NewClient(conn)
	s.conns.Store(conn.RemoteAddr(), client)
	log.Printf("create client id:%v\n", client.GetID())

	// 读取消息内容
	err = prot.IOLoop(client)
	if err != nil {
		log.Printf("[Handle] err:%v", err)
	}

}
