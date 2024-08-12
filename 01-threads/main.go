package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var INT_MAX = 100_000_000
var total_prime_numbers int32 = 1
var cur int32 = 2
var partitions = 10

func checkIfPrime(n int) {
	if n&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		if n%i == 0 {
			return
		}
	}

	atomic.AddInt32(&total_prime_numbers, 1)
}

func singleThreaded() {
	for i := 3; i < INT_MAX; i++ {
		checkIfPrime(i)
	}
}

func batchProcess(wg *sync.WaitGroup, id, l, r int) {
	defer wg.Done()

	start := time.Now()
	for i := l; i < r; i++ {
		checkIfPrime(i)
	}

	fmt.Printf("thread %d completed in [%d %d] %s\n", id, l, r, time.Since(start))
}

func blindMultiThreaded() {
	size := INT_MAX / partitions
	start := 3

	var wg sync.WaitGroup
	for i := 0; i < partitions-1; i++ {
		wg.Add(1)
		go batchProcess(&wg, i, start, start+size)
		start = start + size
	}
	wg.Add(1)
	go batchProcess(&wg, partitions-1, start, INT_MAX)

	wg.Wait()
}

func processOneByOne(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	start := time.Now()
	for {
		x := atomic.AddInt32(&cur, 1)
		if x > int32(INT_MAX) {
			break
		}
		checkIfPrime(int(x))
	}
	fmt.Printf("thread %d completed in %s\n", id, time.Since(start))
}

func multiThreaded() {
	var wg sync.WaitGroup

	for i := 0; i < partitions; i++ {
		wg.Add(1)
		go processOneByOne(&wg, i)
	}

	wg.Wait()

}

func main() {
	start := time.Now()

	// singleThreaded()
	// blindMultiThreaded()
	multiThreaded()

	fmt.Printf("checking till %d found %d in %s\n", INT_MAX, total_prime_numbers, time.Since(start))
}
