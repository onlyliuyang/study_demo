package main

import (
	"fmt"
	"time"
)

type Task struct {
	Num int
}

type Job struct {
	Task Task
}

var (
	MaxWorker = 5
	JobQueue  chan Job //工作通道，模拟需要处理的所有工作
)

type Worker struct {
	id         int
	WorkerPool chan chan Job //工作者池（通道的通道），每个元素都是一个JOB通道
	JobChannel chan Job      //工作通道，每个元素是一个job，worker私有的job
	exit       chan bool     //结束信号
}

func NewWorker(workerPool chan chan Job, id int) Worker {
	fmt.Printf("New a worker(%d)\n", id)
	return Worker{
		id:         id,
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		exit:       make(chan bool),
	}
}

//监听任务和结束信号
func (w Worker) Start() {
	go func() {
		for {
			//将当前的任务队列注册到工作池
			w.WorkerPool <- w.JobChannel
			fmt.Println("Register private JobChannel public WorkerPool", w)
			select {
			case job := <-w.JobChannel:
				//收到任务
				fmt.Println("get a job from private w.JobChannel")
				fmt.Println(job)
				time.Sleep(3 * time.Second)
			case <-w.exit:
				fmt.Println("worker exit", w)
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.exit <- true
	}()
}

//排程中心
type Scheduler struct {
	WorkerPool chan chan Job //工作池
	MaxWorkers int           //最大工作者数
	Workers    []*Worker     //worker队列
}

//创建排程中心
func NewScheduler(maxWorkers int) *Scheduler {
	pool := make(chan chan Job, maxWorkers)	//工作池
	return &Scheduler{WorkerPool:pool, MaxWorkers:maxWorkers}
}

//工作池的处始化
func (s *Scheduler) Create() {
	workers := make([]*Worker, s.MaxWorkers)
	for i:=0; i<s.MaxWorkers; i++ {
		worker := NewWorker(s.WorkerPool, i)
		worker.Start()
		workers[i] = &worker
	}

	s.Workers = workers
	go s.schedule()
}

//工作池的关闭
func (s *Scheduler) Shutdown() {
	workers := s.Workers
	for _, w := range workers {
		w.Stop()
	}
	time.Sleep(time.Second)
	close(s.WorkerPool)
}

//排程
func (s *Scheduler) schedule() {
	for {
		select {
		case job := <-JobQueue:
			fmt.Println("get a job from JobQueue")
			go func(job Job) {
				//从workPool获取jobChannel, 忙时阻塞
				jobChannel := <- s.WorkerPool
				fmt.Println(jobChannel)
				fmt.Println("get a private jobChannel from public s.WorkerPool, jobId:", job.Task.Num)
			}(job)
		}
	}
}

func main()  {
	JobQueue = make(chan Job, 5)
	scheduler := NewScheduler(MaxWorker)
	scheduler.Create()
	time.Sleep(1 * time.Second)

	go createJobQueue()
	time.Sleep(100 * time.Second)
	scheduler.Shutdown()
	time.Sleep(10 * time.Second)
}

//模拟创建任务
func createJobQueue()  {
	for i:=0; i<100; i++ {
		task := Task{Num:i}
		job := Job{Task:task}
		JobQueue <- job
		fmt.Println("JobQueue add one job")
		time.Sleep(1 * time.Second)
	}
}