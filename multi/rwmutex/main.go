package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.RWMutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string, id int) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.RLock()
	defer fmt.Println("Value unlock.")
	defer c.mux.RUnlock()
	v := c.v[key]
	fmt.Println("v is ", v)
	return v
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 10; i++ {
		go c.Inc("somekey", i)
	}
	go c.Value("somekey")
	time.Sleep(time.Second)
}
