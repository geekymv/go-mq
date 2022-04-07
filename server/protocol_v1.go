package server

import (
	"bytes"
	"errors"
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

/*
SUB <topic_name> <channel_name>
*/
func (p *protocolV1) SUB(client *clientV1, params [][]byte) ([]byte, error) {
	if len(params) < 3 {
		return nil, errors.New("params len invalid")
	}
	topicName := params[1]
	channelName := params[2]

	// 获取 topic 和 channel
	t := p.server.GetTopic(string(topicName))
	channel := t.GetChannel(string(channelName))

	// 将 client 添加到 channel
	if err := channel.AddClient(client.ID, client); err != nil {
		return nil, err
	}

	// 将 channel 和 client 关联
	client.Channel = channel

	return []byte("OK"), nil
}
