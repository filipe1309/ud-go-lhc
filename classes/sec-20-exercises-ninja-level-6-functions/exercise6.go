package main

/**
 * Bulid and use an anonymous func
 */

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Anonymous function")
	}()
}
