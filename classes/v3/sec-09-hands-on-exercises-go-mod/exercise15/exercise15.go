package main

/**
 * Using the code you wrote in the previous hands-on exercise:
 *	- go mod init exercise15
 *	- Use a function from the package found at github.com/GoesToEleven/puppy,
 *	but make your code depend on v1.2.0
 *		- go get github.com/GoesToEleven/puppy@v1.2.0
 *	- Inspect your go.mod file
 */

import (
	"fmt"
	"github.com/GoesToEleven/puppy"
)

func main() {
	fmt.Println(puppy.Barks())
	puppy.From12()
	fmt.Println("Hello, exercise 15")
}
