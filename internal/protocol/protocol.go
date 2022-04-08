package protocol

import (
	"encoding/binary"
	"io"
	"net"
)

const MagicV1 = "v1"

type Client interface {
	Close() error
	GetID() int64
}

type Protocol interface {
	NewClient(net.Conn) Client
	IOLoop(Client) error
}

/*
SendFramedResponse
size 			 frameType  data
[4-byte data length] [4-byte]  [N-byte]

size 存储frameType字节数 + data字节数
*/
func SendFramedResponse(w io.Writer, frameType int32, data []byte) (int, error) {
	buf := make([]byte, 4)

	// 发送数据包大小
	size := 4 + uint32(len(data))
	binary.BigEndian.PutUint32(buf, size)
	n, err := w.Write(buf)
	if err != nil {
		return n, err
	}

	// 发送 frameType
	binary.BigEndian.PutUint32(buf, uint32(frameType))
	n, err = w.Write(buf)
	if err != nil {
		return n + 4, err
	}
	// 发送 data
	n, err = w.Write(data)

	return n + 8, err
}
