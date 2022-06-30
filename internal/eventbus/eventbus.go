package eventbus

import "errors"

var ErrTopicNotFound = errors.New("topic not found")
var ErrTopicAlredyExists = errors.New("topic alredy exists")

type Topic string

type event any

type EventBus interface {
	Subscribe(t Topic) (Subscriber, error)
	Publish(t Topic, e event) error
	CreateTopic(t Topic) error
}

type eventBus struct {
	subscribers map[Topic][]chan any
}

func NewEventBus() EventBus {
	return &eventBus{
		subscribers: make(map[Topic][]chan any),
	}
}

var _ EventBus = (*eventBus)(nil)

func (b *eventBus) CreateTopic(t Topic) error {
	_, ok := b.subscribers[t]
	if ok {
		return ErrTopicAlredyExists
	}

	return nil
}

func (b *eventBus) Subscribe(t Topic) (Subscriber, error) {
	_, ok := b.subscribers[t]
	if !ok {
		return nil, ErrTopicNotFound
	}

	newchannel := make(chan any)
	subscriber := NewSubscriber(newchannel)
	b.subscribers[t] = append(b.subscribers[t], newchannel)

	return subscriber, nil
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
