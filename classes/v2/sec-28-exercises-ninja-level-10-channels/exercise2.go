package main

/**
 * Get this code running:
 * https://play.golang.org/p/oB-p3KMiH6
 * solution: https://play.golang.org/p/oB-p3KMiH6
 * https://play.golang.org/p/_DBRueImEq
 * solution: https://play.golang.org/p/SHr3lpX4so
 */

import (
	"fmt"
)

func main() {
	cs := make(chan int)
	// cs := make(chan<- int) // send only channel
	// cr := make(<-chan int) // receive only channel

	go func() {
		cs <- 42
	}()
	fmt.Println("cs:\t", <-cs)

	fmt.Println("-----")
	fmt.Printf("cs\t%T\n", cs)
}
