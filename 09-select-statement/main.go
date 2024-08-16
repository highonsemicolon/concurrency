package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from channel 2"
	}()

	for range 2 {
		select {
		case m := <-ch1:
			fmt.Println(m)

		case m := <-ch2:
			fmt.Println(m)
		}
	}

}
