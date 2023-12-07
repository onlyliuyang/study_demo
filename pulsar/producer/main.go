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
		fmt.Println("Pulsar connect fail: ", err.Error())
		return
	}
	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})

	for i := 0; i < 100; i++ {
		_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("This is my %d message", i)),
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
	defer producer.Close()

	if err != nil {
		fmt.Println("Failed to publish message", err)
		return
	}
	fmt.Println("Published message")
}
