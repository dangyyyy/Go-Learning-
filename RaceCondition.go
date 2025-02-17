package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	count int
	mu    *sync.RWMutex
}

func (c *counter) inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}
func (c *counter) value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func main() {
	c := counter{
		mu: new(sync.RWMutex),
	}
	for i := 0; i < 1000; i++ {

		go func() {
			c.inc()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(c.value())
}
