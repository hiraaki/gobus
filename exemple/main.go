package main

import (
	"fmt"
	"gobus/internal/group"
	gobus "gobus/pkg"
	"log"
	"time"
)

func Dogs() {
	err := gobus.NewTopic("dog")
	if err != nil {
		panic(err)
	}

	subscriber, err := gobus.Subscribe("dog")
	if err != nil {
		panic(err)
	}

	go func() {
		for msg := range subscriber.Channel() {
			fmt.Println(msg)
		}
	}()

	err = gobus.Publish("dog", "german shepherd")
	if err != nil {
		log.Println(err)
	}
}

func Conssuergroups() {
	c := make(chan any)

	r := make(chan any)

	consumer := group.NewConsumer(r)

	cg := group.NewConsumerGroup(c)

	cg.AddConsumer(consumer)

	fmt.Println(cg.Size())

	go func() {
		c <- "golden"
	}()

	go func() {
		consumer, err := cg.Consumer(0)
		if err != nil {
			panic(err)
		}

		event, ok := <-consumer.Thread()
		if !ok {
			fmt.Println("clesed thread")
			return
		}

		fmt.Println("thread response: ", event)

		consumer.Close()
	}()

	<-time.After(time.Second * 5)
	fmt.Println(cg.Size())
}

func main() {
	Conssuergroups()
}
