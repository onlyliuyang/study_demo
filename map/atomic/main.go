package main

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"
)

func KeyNamed(key string) string {
	return strings.ToLower(key)
}

type Map struct {
	Value atomic.Value
}

func (m *Map) Store(values map[string]*atomic.Value) {
	dst := make(map[string]*atomic.Value, len(values))
	for k, v := range values {
		dst[k] = v
	}
	m.Value.Store(dst)
}

func (m *Map) Load() map[string]*atomic.Value {
	src := m.Value.Load().(map[string]*atomic.Value)
	dst := make(map[string]*atomic.Value, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func (m *Map) Exists(key string) bool {
	_, ok := m.Load()[KeyNamed(key)]
	return ok
}

func (m *Map) Get(key string) *atomic.Value {
	v, ok := m.Load()[KeyNamed(key)]
	if ok {
		return v
	}
	return &atomic.Value{}
}

func (m *Map) Keys() []string {
	values := m.Load()
	keys := make([]string, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	return keys
}

func main()  {
	groutines := 2
	m := &Map{}

	for i:=0; i<groutines; i++ {
		go func(i int) {
			for i:=0; i<100; i++ {
				var v atomic.Value
				key := fmt.Sprintf("key_%d", i)
				value := make(map[string]*atomic.Value)
				value[key] = &v
				v.Store(i)
				m.Store(value)
				//fmt.Println(i, value)
			}
		}(i)

		go func(i int) {
			for i:=0; i<100; i++ {
				key := fmt.Sprintf("key_%d", i)

				fmt.Println("key =>", key,  m.Exists(key))
				fmt.Println("keys =>", m.Keys())
				fmt.Println("load: ", m.Load())
			}
		}(i)
	}

	time.Sleep(10 * time.Second)
}