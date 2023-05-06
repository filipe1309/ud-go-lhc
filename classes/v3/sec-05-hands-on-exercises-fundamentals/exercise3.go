package main

/**
 * We learned about bitwize operations in the last section.
 * Now let's learn about "iota" and use it in a program
 * - To learn about iota
 *   - https://golang.org/ref/spec#Iota
 *   - https://golang.org/doc/effective_go#constants
 * - modify this program to use iota
 *  - https://go.dev/play/p/jZUmqlhqalC
 *  - note how iota only need to be declared once
 */

import "fmt"

// const (
// 	c0 = iota // 0
// 	c1 = iota // 1
// 	c2 = iota // 2
// )

// const (
// 	c3 = iota // 0
// 	c4        // 1
// 	c5        // 2
// 	c6        // 3
// )

const (
	_  = iota // 0
	c1        // 1
	c2        // 2
	c3        // 3
	c4        // 4
	c5        // 5
	c6        // 6
)

func main() {
	fmt.Println(c1, c2)
	fmt.Println(c3, c4, c5, c6)
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<c1, 1<<c1)
	fmt.Printf("%d \t %b\n", 1<<c2, 1<<c2)
	fmt.Printf("%d \t %b\n", 1<<c3, 1<<c3)
	fmt.Printf("%d \t %b\n", 1<<c4, 1<<c4)
	fmt.Printf("%d \t %b\n", 1<<c5, 1<<c5)
	fmt.Printf("%d \t %b\n", 1<<c6, 1<<c6)
}
