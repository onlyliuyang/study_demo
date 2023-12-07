package queue

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPriorityQueue(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//priorityQueue := NewPriorityQueue()
	rand.Seed(time.Now().Unix())

	t.Log("begin test")

	//随机生成优先级队列
	for i := 0; i < 1000; i++ {
		a := rand.Intn(10)
		rand.Seed(time.Now().UnixNano())
		t.Log(a)
		//go func(i int) {
		//	priorityQueue.Push(func() {
		//		fmt.Println("推送任务的编号为: ", i)
		//		fmt.Println("推送任务的优先级为: ", a)
		//	}, a)
		//}(a)
	}

	//priorityQueue.Consume()
}
