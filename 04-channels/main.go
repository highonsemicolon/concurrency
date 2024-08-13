package main

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"
)

var INT_MAX = 100_000_000
var total_prime_numbers int32 = 1
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

func batchProcess2(ch chan int, done chan struct{}) {
	fmt.Println("here....")
	for i := range ch {
		checkIfPrime(i)
	}

	fmt.Printf("thread completed %v\n", ch)
	done <- struct{}{}
}

func channelBasedMultiThreaded() {
	ch := make(chan int, 1000)
	done := make(chan struct{})

	for i := 0; i < partitions; i++ {
		go batchProcess2(ch, done)
	}

	go func() {
		for i := 3; i < INT_MAX; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := 0; i < partitions; i++ {
		<-done
	}

}

func main() {
	start := time.Now()
	channelBasedMultiThreaded()

	fmt.Printf("checking till %d found %d in %s\n", INT_MAX, total_prime_numbers, time.Since(start))
}
