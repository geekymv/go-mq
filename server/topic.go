package server

import (
	"encoding/binary"
	"log"
)

type Topic struct {
	name       string
	channelMap map[string]*Channel
	// 消息 channel
	memoryMsgChan chan *Message
	server        *MQServer
	startChan     chan int
	idFactory     *ID
}

func NewTopic(topicName string, s *MQServer) *Topic {
	t := &Topic{
		name:          topicName,
		channelMap:    make(map[string]*Channel),
		memoryMsgChan: make(chan *Message, 10000),
		server:        s,
		startChan:     make(chan int, 1),
		idFactory:     NewID(s.NodeID),
	}

	go t.messagePump()

	return t
}

func (t *Topic) GetChannel(channelName string) *Channel {
	channel, ok := t.channelMap[channelName]
	if ok {
		return channel
	}

	// channel 不存在，创建 channel
	channel = NewChannel(t.name, channelName, t.server)
	t.channelMap[channelName] = channel

	return channel
}

func (t *Topic) Start() {
	select {
	case t.startChan <- 1:
	default:
	}
}

// GenerateID 生成 MessageID
func (t *Topic) GenerateID() MessageID {
	var msgId MessageID

	nextID, err := t.idFactory.NextID()
	if err != nil {
		log.Printf("[GenerateID] err:%v\n", err)
	}
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, nextID)
	copy(msgId[:], buf)

	return msgId
}

func (t *Topic) PutMessage(m *Message) error {
	select {
	case t.memoryMsgChan <- m:
	default:
		// TODO write message to backend
	}
	return nil
}

// 将消息分发给 topic 关联的 channel
func (t *Topic) messagePump() {
	var msg *Message

	for {
		select {
		case msg = <-t.memoryMsgChan:
		default:
		}

		// 遍历 channel
		for _, c := range t.channelMap {
			c.PutMessage(msg)
		}
	}

}
