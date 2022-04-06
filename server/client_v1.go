package server

import (
	"net"

	"github.com/geekymv/go-mq/internal/protocol"
)

type clientV1 struct {
	id     int64
	Conn   net.Conn
	server *MQServer
}

func newClientV2(id int64, conn net.Conn, server *MQServer) protocol.Client {
	c := &clientV1{
		id:     id,
		Conn:   conn,
		server: server,
	}
	return c
}

func (c *clientV1) ID() int64 {
	return c.id
}

func (c *clientV1) Close() error {
	return nil
}
