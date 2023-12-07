package lock

import (
	"fmt"
	"sync"
)

type LockMap struct {
	m map[string]interface{}
	sync.Mutex
}

func (l *LockMap) Get(key string) (value interface{}, ok bool) {
	l.Lock()
	defer l.Unlock()

	if value, ok := l.m[key]; ok {
		return value, ok
	}
	return nil, false
}

func (l *LockMap) Set(key string, value interface{}) {
	l.Lock()
	defer l.Unlock()

	l.m[key] = value
}

func (l *LockMap) Delete(key string) {
	l.Lock()
	defer l.Unlock()

	delete(l.m, key)
}

func main()  {
	m := LockMap{
		m:make(map[string]interface{}),
	}

	wg := sync.WaitGroup{}
	for i:=0; i<100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			key := fmt.Sprintf("key_%d", i)
			value := i
			m.Set(key, value)
		}(i)
	}
	wg.Wait()

	//fmt.Println(m)
	for key, value := range m.m {
		fmt.Printf("key => %s, value => %v\n", key, value)
	}
}