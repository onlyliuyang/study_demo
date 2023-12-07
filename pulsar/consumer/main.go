package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"time"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  3 * time.Second,
		ConnectionTimeout: 3 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "after_article_create",
		SubscriptionName: "my-sub",
		Type:             pulsar.Shared,
	})
	defer consumer.Close()

	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received message msgId: %#v --content: '%s'\n", msg.ID(), string(msg.Payload()))
	}
}
