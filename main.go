package main

import (
	"fmt"
	gobus "gobus/pkg"
	"log"
	"time"
)

func main() {
	subscriber := gobus.Subscribe("dog")
	go func() {
		for msg := range subscriber.Channel() {
			fmt.Println(msg)
		}
	}()
	err := gobus.Publish("dog", "german shepherd")
	if err != nil {
		log.Println(err)
	}
	time.Sleep(time.Minute)
}
