package server

import (
	"bufio"
	"net"

	"github.com/geekymv/go-mq/internal/protocol"
)

const defaultBufferSize = 16 * 1024

type clientV1 struct {
	ID int64
	net.Conn
	server *MQServer

	Reader *bufio.Reader
	Writer *bufio.Writer
}

func newClientV1(id int64, conn net.Conn, server *MQServer) protocol.Client {
	c := &clientV1{
		ID:     id,
		Conn:   conn,
		server: server,

		Reader: bufio.NewReaderSize(conn, defaultBufferSize),
		Writer: bufio.NewWriterSize(conn, defaultBufferSize),
	}
	return c
}

func (c *clientV1) GetID() int64 {
	return c.ID
}
