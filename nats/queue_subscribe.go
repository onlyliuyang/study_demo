package main

import (
	"fmt"
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

	nc.QueueSubscribe("foo", "queue.foo", func(msg *nats.Msg) {
		log.Println("Subscribe 1:", string(msg.Data))
	})

	nc.QueueSubscribe("foo", "queue.foo", func(msg *nats.Msg) {
		log.Println("Subscribe 2:", string(msg.Data))
	})

	nc.QueueSubscribe("foo", "queue.foo", func(msg *nats.Msg) {
		log.Println("Subscribe 3:", string(msg.Data))
	})

	for i:=1; i <= 10; i++{
		message := fmt.Sprintf("Message %d", i)

		if err := nc.Publish("foo", []byte(message)); err != nil {
			log.Fatal(err)
		}
	}

	time.Sleep(2 * time.Second)
}