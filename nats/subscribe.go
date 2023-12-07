package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Println("Subscribe 1:", string(msg.Data))
	})

	nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Println("Subscribe 2:", string(msg.Data))
	})

	nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Println("Subscribe 3:", string(msg.Data))
	})

	if err := nc.Publish("foo", []byte("Here's some stuff")); err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)
}