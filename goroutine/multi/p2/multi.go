package main

import "sync"

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
