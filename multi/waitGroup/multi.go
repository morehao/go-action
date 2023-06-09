package main

import "sync"

type Multi struct {
	wg *sync.WaitGroup
	ch chan struct{}
}

func NewMulti(multiNum int) *Multi {
	return &Multi{
		wg: &sync.WaitGroup{},
		ch: make(chan struct{}, multiNum),
	}
}

func (m *Multi) Add(num int) {
	m.wg.Add(num)
	for i := 0; i < num; i++ {
		m.ch <- struct{}{}
	}
}

func (m *Multi) Done() {
	<-m.ch
	m.wg.Done()
}

func (m *Multi) Wait() {
	m.wg.Wait()
}

// func (m *Multi) Run(f func()) {
// 	defer m.Done()
// 	m.Add(1)
// 	go f()
// }
