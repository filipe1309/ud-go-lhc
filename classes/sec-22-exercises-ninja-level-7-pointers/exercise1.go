package main

/*
 * Create a value and assign it to a variable.
 * Print the address of that value.
 */

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
}
