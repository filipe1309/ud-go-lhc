package main

/**
 * Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
 */

import (
	"fmt"
)

func main() {
	defer myDefer()
	fmt.Println("Hello, playground")
	fmt.Println("Hello, playground 2")
	foo()
}

func myDefer() {
	fmt.Println("Defered function")
}

func foo() {
	defer fmt.Println("Defered FOO function")
	fmt.Println("FOO function")
}
