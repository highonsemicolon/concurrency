package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	version int64
	value   int64
}

func (c *Counter) increment() bool {
	oldVersion := atomic.LoadInt64(&c.version)
	oldValue := atomic.LoadInt64(&c.value)

	newValue := oldValue + 1

	if atomic.CompareAndSwapInt64(&c.version, oldVersion, oldVersion+1) {
		atomic.StoreInt64(&c.value, newValue)
		return true
	}
	return false
}

func worker(wg *sync.WaitGroup, c *Counter) {
	defer wg.Done()

	success := c.increment()
	if success {
		//fmt.Println("increment succeeded")
	} else {
		fmt.Println("increment failed due to version mismatch")
	}

}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{value: 0, version: 0}

	for range 200 {
		wg.Add(1)
		go worker(&wg, counter)
	}

	wg.Wait()
	fmt.Printf("final value: %d", counter.value)
}
