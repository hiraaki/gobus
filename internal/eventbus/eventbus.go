package eventbus

import "errors"

var ErrTopicNotFound = errors.New("topic not found")

type Topic string

type event interface{}

type EventBus interface {
	Subscribe(t Topic) Subscriber
	Publish(t Topic, e event) error
}

type eventBus struct {
	subscribers map[Topic][]chan interface{}
}

func NewEventBus() EventBus {
	return &eventBus{
		subscribers: make(map[Topic][]chan interface{}),
	}
}

var _ EventBus = (*eventBus)(nil)

func (b *eventBus) Subscribe(t Topic) Subscriber {
	newchannel := make(chan interface{})
	subscriber := NewSubscriber(newchannel)
	b.subscribers[t] = append(b.subscribers[t], newchannel)
	return subscriber
}

func (b *eventBus) Publish(t Topic, e event) error {
	subscribers := b.subscribers[t]
	if subscribers == nil {
		return ErrTopicNotFound
	}
	for _, r := range subscribers {
		r <- e
	}
	return nil
}
