package queue

import (
	"fmt"
	"sync"
)

type PriorityQueue struct {
	mLock      sync.Mutex         //互斥锁，queues和priorities并发操作
	queues     map[int]chan *Task //优先队列map
	pushChan   chan *Task         //推送任务管道
	priorities []int              //记录优先级的切片，优先级从小到大排列
}

type Task struct {
	priority int    //任务的优先级
	f        func() //任务的执行函数
}

func NewPriorityQueue() *PriorityQueue {
	priorityQueue := &PriorityQueue{
		queues:   make(map[int]chan *Task),
		pushChan: make(chan *Task, 100),
	}

	go priorityQueue.ListenPushChan()
	return priorityQueue
}

func (pq *PriorityQueue) ListenPushChan() {
	for {
		select {
		case taskEle := <-pq.pushChan:
			priority := taskEle.priority
			pq.mLock.Lock()
			if v, ok := pq.queues[priority]; ok {
				pq.mLock.Unlock()
				//将推送的任务塞到对应优先级的队列中
				v <- taskEle
				continue
			}

			//如果这是一个新的优先级，则需要插入优先级切片，并且新建一个优先级的queue
			//通过二分法查找
			index := pq.getNewPriorityInsertIndex(priority, 0, len(pq.priorities)-1)
			pq.moveNextPriorities(index, priority)

			//创建一个新的队列
			pq.queues[priority] = make(chan *Task, 1000)

			//将任务塞到新的队列中
			pq.queues[priority] <- taskEle
			pq.mLock.Unlock()
		}
	}
}

// index 右侧元素均需要向后移动一个单位
func (pq *PriorityQueue) moveNextPriorities(index, priority int) {
	pq.priorities = append(pq.priorities, 0)
	copy(pq.priorities[index+1:], pq.priorities[index:])
	pq.priorities[index] = priority
}

func (pq *PriorityQueue) getNewPriorityInsertIndex(priority int, leftIndex, rightIndex int) (index int) {
	//如果当前切片没有元素，则插入的index就是0
	if len(pq.priorities) == 0 {
		return 0
	}

	length := rightIndex - leftIndex
	//如果当前切片中最小的元素都超过了插入的优先级，则插入位置应该是最左边
	if pq.priorities[leftIndex] >= priority {
		return leftIndex
	}

	//如果当前切片中最大的元素都没超过插入的优先级，则插入的公交地图应该是最右边
	if pq.priorities[rightIndex] <= priority {
		return rightIndex + 1
	}

	if length == 1 && pq.priorities[leftIndex] < priority && pq.priorities[rightIndex] >= priority {
		return leftIndex + 1
	}

	middleVal := pq.priorities[leftIndex+length/2]
	if priority <= middleVal {
		return pq.getNewPriorityInsertIndex(priority, leftIndex, leftIndex+length/2)
	} else {
		return pq.getNewPriorityInsertIndex(priority, leftIndex+length/2, rightIndex)
	}
}

//取出最高优先级队列中的一个任务
func (pq *PriorityQueue) Pop() *Task {
	pq.mLock.Lock()
	defer pq.mLock.Unlock()

	for i := len(pq.priorities) - 1; i >= 0; i-- {
		if len(pq.queues[pq.priorities[i]]) == 0 {
			continue
		}

		//如果当前优先级队列里有任务，则取出一个任务
		return <-pq.queues[pq.priorities[i]]
	}
	return nil
}

//写入任务
func (pq *PriorityQueue) Push(f func(), priority int) {
	fmt.Println("task push")
	pq.pushChan <- &Task{
		f:        f,
		priority: priority,
	}
}

//消费者获取最高优先级的任务
func (pq *PriorityQueue) Consume() {
	for {
		task := pq.Pop()
		//fmt.Println(task)
		//time.Sleep(5 * time.Second)
		if task == nil {
			continue
		}

		//执行任务
		task.f()
	}
}
