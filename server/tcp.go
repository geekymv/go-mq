package server

import (
	"io"
	"log"
	"net"

	"github.com/geekymv/go-mq/internal/protocol"
)

type tcpServer struct {
	server *MQServer
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
	var p protocol.Protocol
	switch v {
	case protocol.MagicV1:
		p = &protocolV1{server: s.server}
	default:
		conn.Close()
		log.Printf("client(%s) error version '%s'\n", conn.RemoteAddr(), v)
	}
	// 创建 client
	client := p.NewClient(conn)
	log.Printf("create client id:%v\n", client.ID())

	// 读取消息内容

}
