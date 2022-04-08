package server

type Consumer interface {
}

type Channel struct {
	topicName string
	name      string
	server    *MQServer

	clients map[int64]Consumer
	// 消息 channel
	memoryMsgChan chan *Message
}

func NewChannel(topicName string, name string, s *MQServer) *Channel {
	c := &Channel{
		topicName:     topicName,
		name:          name,
		server:        s,
		clients:       make(map[int64]Consumer),
		memoryMsgChan: make(chan *Message, 10000),
	}
	return c
}

func (c *Channel) AddClient(clientID int64, client Consumer) error {
	_, ok := c.clients[clientID]
	if ok {
		return nil
	}

	c.clients[clientID] = client

	return nil
}

/*
PutMessage 方法中 memoryMsgChan channel 满了，会执行 default
*/
func (c *Channel) PutMessage(m *Message) error {
	select {
	case c.memoryMsgChan <- m:
	default:
		// TODO write message to backend
	}
	return nil
}
