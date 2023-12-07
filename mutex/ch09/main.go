package main

import (
	"sync"
	"time"
)

type pass struct {
	RWM sync.RWMutex
	pwd string
}

var RoomPass = pass{pwd:"initPass"}

func change(p *pass, pwd string) {
	p.RWM.Lock()
	defer p.RWM.Unlock()

	time.Sleep(2 * time.Second)
	p.pwd = pwd
}

func getPwd(p *pass) string {
	p.RWM.RLock()
	defer p.RWM.RUnlock()

	pwd := p.pwd
	return pwd
}
