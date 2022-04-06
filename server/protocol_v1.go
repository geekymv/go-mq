package server

import (
	"net"
	"sync/atomic"

	"github.com/geekymv/go-mq/internal/protocol"
)

type protocolV1 struct {
	server *MQServer
}

func (p *protocolV1) NewClient(conn net.Conn) protocol.Client {
	clientID := atomic.AddInt64(&p.server.ClientIDSequence, 1)
	c := newClientV2(clientID, conn, p.server)
	return c
}
