package main

/**
 * Create a func which returns a func
 * assign the returned func to a variable
 * call the returned func
 */

import (
	"fmt"
)

func main() {
	myVarFunc := myFunc()
	myVarFunc()
}

func myFunc() func() {
	return func() {
		fmt.Println("My returned function")
	}
}
