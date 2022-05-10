package main

import (
	"fmt"
	gobus "gobus/pkg"
	"log"
)

func main() {
	subscriber := gobus.Subscribe("cachorro")
	subscriber2 := gobus.Subscribe("cachorro")
	go func() {
		err := gobus.Publish("cachorro", "cachorroloco")
		if err != nil {
			log.Println(err)
		}
	}()
	msg := <-subscriber.Channel()
	fmt.Println(msg)
	msg = <-subscriber2.Channel()
	fmt.Println(msg)
}
