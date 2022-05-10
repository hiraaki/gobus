package eventbus

type Subscriber interface {
	Channel() <-chan interface{}
}

type subscriber struct {
	channel chan interface{}
}

func NewSubscriber(channel chan interface{}) Subscriber {
	return &subscriber{
		channel: channel,
	}
}

func (s *subscriber) Channel() <-chan interface{} {
	return s.channel
}
