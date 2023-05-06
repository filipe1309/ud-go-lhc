package main

/**
 * Write a program that uses print verbs to show the following numbers
 * 	- 747
 * 	- 911
 * 	- 90210
 * 	as
 * 	- decimal
 * 	- binary
 * 	- hex
 */

import "fmt"

func main() {
	n1, n2, n3 := 747, 911, 90210
	fmt.Printf("%d\t%b\t\t%#x\n", n1, n1, n1)
	fmt.Printf("%d\t%b\t\t%#x\n", n2, n2, n2)
	fmt.Printf("%d\t%b\t%#x\n", n3, n3, n3)
}
