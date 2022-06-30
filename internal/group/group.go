package group

import (
	"errors"
	"sync"
)

type key uint

var ErrOutOfRange = errors.New("index out of range")

type ConsumerGroup interface {
	Consumer(k key) (Consumer, error)
	Add(c Consumer)
	Size() uint
}

type consumerGroup struct {
	mu        sync.Mutex
	size      key
	t         thread
	consumers map[key]Consumer
}

var _ ConsumerGroup = (*consumerGroup)(nil)

func NewConsumerGroup(t thread) ConsumerGroup {
	cg := consumerGroup{
		size:      0,
		t:         t,
		consumers: make(map[key]Consumer),
	}

	go cg.runner()

	return &cg
}

func (c *consumerGroup) runner() {
	for e := range c.t {
		for k, consumer := range c.consumers {
			_, ok := <-consumer.Done()
			if !ok {
				c.remove(k)
				continue
			}

			consumer.Thread() <- e
		}
	}
	c.shutdown()
}

func (c *consumerGroup) Consumer(k key) (Consumer, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if k >= c.size {
		return nil, ErrOutOfRange
	}
	return c.consumers[k], nil
}

func (c *consumerGroup) Add(consumer Consumer) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.consumers[c.size] = consumer

	c.size++
}

func (c *consumerGroup) Size() uint {
	return uint(c.size)
}

func (c *consumerGroup) shutdown() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, consumer := range c.consumers {
		consumer.Close()
		delete(c.consumers, key)
	}
	c.size = 0
}

func (c *consumerGroup) remove(k key) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.size--
	delete(c.consumers, k)
}
