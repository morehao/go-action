package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := NewMulti(3)
	for i := 0; i < 4; i++ {
		go Read(m, i)
	}
	m.Wait()
}

func Read(m *multi, i int) {
	defer m.Done()
	m.Add(1)
	fmt.Println("Read i", i)
	time.Sleep(2 * time.Second)
}

type multi struct {
	c  chan struct{}
	wg *sync.WaitGroup
}

func NewMulti(maxNum int) *multi {
	return &multi{
		c: make(chan struct{}, maxNum),
		// wg: new(sync.WaitGroup),
		wg: &sync.WaitGroup{},
	}
}

func (m *multi) Add(num int) {
	m.wg.Add(num)
	for i := 0; i < num; i++ {
		m.c <- struct{}{}
	}
}

func (m *multi) Done() {
	<-m.c
	m.wg.Done()
}

func (m *multi) Wait() {
	m.wg.Wait()
}
