package main

import (
	"fmt"
	"time"
)

/*方案二：将方案一封装*/
func main() {
	var m = NewMulti(3)
	userCount := 10
	for i := 0; i < userCount; i++ {
		go Read(m, i)
	}

	m.Wait()
}

func Read(m *multi, i int) {
	defer m.Done()
	m.Add(1)

	fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second)
}
