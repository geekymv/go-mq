package server

type Consumer interface {
}

type Channel struct {
	topicName string
	name      string
	server    *MQServer

	clients map[int64]Consumer
}

func NewChannel(topicName string, name string, s *MQServer) *Channel {
	c := &Channel{
		topicName: topicName,
		name:      name,
		server:    s,
		clients:   make(map[int64]Consumer),
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
