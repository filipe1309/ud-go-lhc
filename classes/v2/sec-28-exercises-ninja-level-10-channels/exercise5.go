package main

/**
 * Show the comma ok idiom starting with this code (https://play.golang.org/p/YHOMV9NYKK).
 */

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
