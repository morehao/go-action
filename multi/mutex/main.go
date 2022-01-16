package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string, id int) {
	c.mux.Lock()
	fmt.Printf("%d. Inc lock.\n", id)
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
	fmt.Printf("%d. Inc unlock.\n", id)
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	fmt.Println("Value lock.")
	// Lock so only one goroutine at a time can access the map c.v.
	defer fmt.Println("Value unlock.")
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 10; i++ {
		go c.Inc("somekey", i)
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
