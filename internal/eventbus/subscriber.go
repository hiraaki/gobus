package eventbus

type Subscriber interface {
	Channel() <-chan any
	Close()
}

type subscriber struct {
	channel chan any
}

func NewSubscriber(channel chan any) Subscriber {
	return &subscriber{
		channel: channel,
	}
}

func (s *subscriber) Channel() <-chan any {
	return s.channel
}

func (s *subscriber) Close() {
	close(s.channel)
}
