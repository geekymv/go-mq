package server

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
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
PUB <topic_name>\n
[ 4-byte size in bytes ][ N-byte binary data ]
*/
func (p *protocolV1) PUB(client *clientV1, params [][]byte) ([]byte, error) {
	topicName := string(params[1])

	// 读取消息内容长度
	buf := make([]byte, 4)
	_, err := io.ReadFull(client.Reader, buf)
	if err != nil {
		return nil, err
	}
	bodyLen := binary.BigEndian.Uint32(buf)
	messageBody := make([]byte, bodyLen)
	// 读取消息内容
	_, err = io.ReadFull(client.Reader, messageBody)
	if err != nil {
		return nil, err
	}

	// 获取 topic，并发送消息
	topic := p.server.GetTopic(topicName)
	msg := NewMessage(topic.GenerateID(), messageBody)
	err = topic.PutMessage(msg)

	return []byte("OK"), err
}

/*
SUB <topic_name> <channel_name>\n
*/
func (p *protocolV1) SUB(client *clientV1, params [][]byte) ([]byte, error) {
	if len(params) < 3 {
		return nil, errors.New("params len invalid")
	}
	topicName := string(params[1])
	channelName := string(params[2])
	// TODO 验证 topicName channelName 合法性

	// 获取 topic 和 channel
	t := p.server.GetTopic(topicName)
	channel := t.GetChannel(channelName)

	// 将 client 添加到 channel
	if err := channel.AddClient(client.ID, client); err != nil {
		return nil, err
	}

	// 将 channel 和 client 关联
	client.Channel = channel

	return []byte("OK"), nil
}
