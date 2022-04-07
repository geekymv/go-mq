package protocol

import (
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
