package main

import (
	"fmt"
	"sync"
	"time"
)

type Job func()

type Pool struct {
	workQueue chan Job
	wg        sync.WaitGroup
}

func newPool(n int) *Pool {
	pool := &Pool{
		workQueue: make(chan Job),
	}

	pool.wg.Add(n)
	for range n {
		go pool.worker()
	}

	return pool
}

func (p *Pool) worker() {
	defer p.wg.Done()
	for job := range p.workQueue {
		job()
	}
}

func (p *Pool) add(job Job) {
	p.workQueue <- job
}

func (p *Pool) wait() {
	close(p.workQueue)
	p.wg.Wait()
}

func main() {

	pool := newPool(10)
	for i := range 500 {
		job := func() {
			time.Sleep(1 * time.Second)
			fmt.Println("job completed", i)
		}

		pool.add(job)
	}

	pool.wait()
}
