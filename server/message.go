package server

import "time"

const (
	MsgIDLength       = 8
	minValidMsgLength = MsgIDLength + 8 + 2
)

type MessageID [MsgIDLength]byte

type Message struct {
	ID        MessageID
	Body      []byte
	Timestamp int64
	ClientID  int64
}

func NewMessage(id MessageID, body []byte) *Message {
	return &Message{
		ID:        id,
		Body:      body,
		Timestamp: time.Now().UnixNano(),
	}
}
