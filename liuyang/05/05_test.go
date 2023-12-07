package main

import (
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log("begin", time.Now().Format("2006-01-02_15:04:05"))
		time.Sleep(1 * time.Second)
		t.Log("end", time.Now().Format("2006-01-02_15:04:05"))
	}
}

func TestTick(t *testing.T) {
	t1 := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t1.C:
			t.Log("begin", time.Now().Format("2006-01-02_15:04:05"))
			t.Log("Do something 1s")
			time.Sleep(time.Second * 1)
			t.Log("end", time.Now().Format("2006-01-02_15:04:05"))
		}
	}
}
