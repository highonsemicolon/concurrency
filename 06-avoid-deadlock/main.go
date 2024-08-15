package main

import (
	"fmt"
	"sync"
)

var mu1, mu2 sync.Mutex

func acquireLock(id, lockId int, mu *sync.Mutex) {
	fmt.Printf("goroutine %d wants to lock %d\n", id, lockId)
	mu.Lock()
	fmt.Printf("goroutine %d acquired lock %d\n", id, lockId)
}

func releaseLock(id, lockId int, mu *sync.Mutex) {
	mu.Unlock()
	fmt.Printf("goroutine %d released lock %d\n", id, lockId)
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		acquireLock(1, 1, &mu1)
		acquireLock(1, 2, &mu2)

		releaseLock(1, 1, &mu1)
		releaseLock(1, 2, &mu2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		acquireLock(2, 1, &mu1)
		acquireLock(2, 2, &mu2)

		releaseLock(2, 1, &mu1)
		releaseLock(2, 2, &mu2)
	}()

	wg.Wait()
	fmt.Println("done!")

}
