package main

/**
 * Get this code running: https://play.golang.org/p/j-EA6003P0
 * using func literal, aka, anonymous self-executing func
 * solution: https://play.golang.org/p/SHr3lpX4so
 * using a buffered channel
 * solution: https://play.golang.org/p/Y0Hx6IZc3U
 */

import (
	"fmt"
)

func main() {
	c := make(chan int)     // unbuffered channel
	c2 := make(chan int, 1) // buffered channel

	// Func literal, aka, anonymous self-executing func
	go func() {
		c <- 42
	}()
	fmt.Println("Func Literal:\t ", <-c)

	// Buffered channel
	c2 <- 42
	fmt.Println("Buffered channel:", <-c2)
}
