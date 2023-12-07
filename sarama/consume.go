package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"time"
)

//消费者
func SarmaConsumer()  {
	var wg sync.WaitGroup
	consumer, err := sarama.NewConsumer([]string{"kafka0:9092", "kafka1:9093", "kafka2:9094"}, nil)
	if err != nil {
		fmt.Println("Failed to start consumer: ", err.Error())
		return
	}

	//通过topic获取所有分区
	partitionList, err := consumer.Partitions("go_server_topic")
	if err != nil {
		fmt.Println("Failed to get the list of partition: ", err.Error())
		return
	}

	fmt.Println(partitionList)

	for partition := range partitionList {
		//针对每个分区创建一个分给消费都
		pc, err := consumer.ConsumePartition("go_server_topic", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}

		wg.Add(1)
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range partitionConsumer.Messages() {
				fmt.Printf("Partition: %d, offset: %d, key: %s, value: %s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}

//生产者
func SaramaProduce()  {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"kafka0:9092", "kafka1:9093", "kafka2:9094"}, config)
	if err != nil {
		fmt.Println("producer closed, err: ", err)
		return
	}
	defer client.Close()

	var i int
	for {
		//构造消息
		value := fmt.Sprintf("producer kafka messages..., ID:%d", i)
		msg := &sarama.ProducerMessage{
			Topic:     "go_server_topic",
			Key:       nil,
			Value:     sarama.StringEncoder(value),
			Headers:   nil,
			Metadata:  nil,
			Offset:    0,
			Partition: 0,
			Timestamp: time.Time{},
		}

		//发送消息
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send msg failed, err: ", err)
		}
		fmt.Printf("pid: %v, offset:%v\n", pid, offset)
	}

}

func main()  {
	var wg sync.WaitGroup
	wg.Add(2)
	go SarmaConsumer()
	go SaramaProduce()

	wg.Wait()
}