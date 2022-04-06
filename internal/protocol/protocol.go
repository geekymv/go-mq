package protocol

import (
	"net"
)

const MagicV1 = "v1"

type Client interface {
	Close() error
	ID() int64
}

type Protocol interface {
	NewClient(conn net.Conn) Client
}
