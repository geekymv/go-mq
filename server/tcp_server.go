package server

import (
	"log"
	"net"
)

type Handler interface {
	Handle(conn net.Conn)
}

func TCPServer(listener net.Listener, handler Handler) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok {
				log.Printf("accept error:%v", nerr)
				continue
			}
			break
		}
		go handler.Handle(conn)
	}

	return nil
}
