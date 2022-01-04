package main

import "sync"

type multi1 struct {
	c  chan struct{}
	wg *sync.WaitGroup
}

func NewMulti1(multiNum int) *multi1 {
	return &multi1{
		c:  make(chan struct{}, multiNum),
		wg: new(sync.WaitGroup),
	}
}

func (m *multi1) Add(delta int) {
	m.wg.Add(delta)
	for i := 0; i < delta; i++ {
		m.c <- struct{}{}
	}
}

func (m *multi1) Done() {
	<-m.c
	m.wg.Done()
}

func (m *multi1) Wait() {
	m.wg.Wait()
}
