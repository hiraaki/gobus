package gobus

import "gobus/internal/eventbus"

func Publish(t eventbus.Topic, event interface{}) error {
	return instance().Publish(t, event)
}

func Subscribe(t eventbus.Topic) (eventbus.Subscriber, error) {
	return instance().Subscribe(t)
}

func NewTopic(t eventbus.Topic) error {
	return instance().CreateTopic(t)
}
