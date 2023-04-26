package main

/**
 * Write a program that
 * - launches 10 goroutines
 * - each goroutine adds 10 numbers to a channel
 * - pull the numbers off the channel and print them
 */

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			for k := 0; k < 10; k++ {
				c <- k
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
