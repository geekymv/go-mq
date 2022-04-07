package server

type Topic struct {
	name       string
	channelMap map[string]*Channel
	server     *MQServer
	startChan  chan int
}

func NewTopic(topicName string, s *MQServer) *Topic {
	t := &Topic{
		name:       topicName,
		channelMap: make(map[string]*Channel),
		server:     s,
		startChan:  make(chan int, 1),
	}
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
