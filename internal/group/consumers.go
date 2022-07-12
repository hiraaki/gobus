package group

import (
	"fmt"
	"sync"
)

type thread chan any

type Consumer interface {
	Thread() thread
	Close()
	Done() <-chan struct{}
}

type consumer struct {
	once   sync.Once
	done   chan struct{}
	thread thread
}

var _ Consumer = (*consumer)(nil)

func NewConsumer(t thread) Consumer {
	c := consumer{
		done:   make(chan struct{}),
		thread: t,
	}

	return &c
}

func (c *consumer) Thread() thread {
	return c.thread
}

func (c *consumer) Close() {
	c.once.Do(func() {
		fmt.Println("vai que é tua")
		close(c.thread)
		close(c.done)
	})
}

func (c *consumer) Done() <-chan struct{} {
	return c.done
}
