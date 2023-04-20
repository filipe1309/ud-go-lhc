package main

/**
 * Assign a func to a variable, then call that func
 */

import (
	"fmt"
)

// var myVarFunc func()

func main() {
	myAnonymousFunc := func() {
		fmt.Println("Anonymous function")
	}
	myVarFunc := myFunc

	myAnonymousFunc()
	myVarFunc()
}

func myFunc() {
	fmt.Println("My function")
}
