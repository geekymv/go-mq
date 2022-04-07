package server

import (
	"bytes"
	"fmt"
	"net"
	"sync/atomic"

	"github.com/geekymv/go-mq/internal/protocol"
)

type protocolV1 struct {
	server *MQServer
}

func (p *protocolV1) NewClient(conn net.Conn) protocol.Client {
	clientID := atomic.AddInt64(&p.server.ClientIDSequence, 1)
	c := newClientV1(clientID, conn, p.server)
	return c
}

func (p *protocolV1) IOLoop(c protocol.Client) error {
	var err error
	var line []byte

	client := c.(*clientV1)
	for {
		line, err = client.Reader.ReadSlice('\n')
		if err != nil {
			break
		}

		// 去除 \n
		line = line[:len(line)-1]
		if len(line) > 0 && line[len(line)-1] == '\r' {
			// 去除 \r
			line = line[:len(line)-1]
		}
		// [][]byte
		params := bytes.Split(line, []byte(" "))
		fmt.Println("params", params)

	}

	return err
}
