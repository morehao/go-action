package main

import (
	"fmt"
	"sync"
	"time"
)

type multi struct {
	c  chan struct{}
	wg *sync.WaitGroup
}

func NewMulti(multiNum int) *multi {
	return &multi{
		c:  make(chan struct{}, multiNum),
		wg: new(sync.WaitGroup),
	}
}

func (m *multi) Add(delta int) {
	m.wg.Add(delta)
	for i := 0; i < delta; i++ {
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

/*方案二：将方案一封装*/
var m = NewMulti(3)

func main() {
	userCount := 10
	for i := 0; i < userCount; i++ {
		go Read(i)
	}

	m.Wait()
}

func Read(i int) {
	defer m.Done()
	m.Add(1)

	fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second)
}
