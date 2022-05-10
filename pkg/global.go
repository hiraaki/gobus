package gobus

import (
	"gobus/internal/eventbus"
	"sync"
)

var globalEventBus eventbus.EventBus
var once sync.Once

func instance() eventbus.EventBus {
	once.Do(func() {
		globalEventBus = eventbus.NewEventBus()
	})
	return globalEventBus
}
