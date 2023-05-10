package main

/**
 * Using the code you wrote in the previous hands-on exercise:
 *	- go mod init exercise14
 *	- Use a function from the package found at github.com/GoesToEleven/puppy
 *		- go get github.com/GoesToEleven/puppy
 *	- Inspect your go.mod file
 *	- Run go mod tidy
 *	- What does go mod tidy do?
 *		- https://go.dev/ref/mod#go-mod-tidy
 */

import (
	"fmt"
	"github.com/GoesToEleven/puppy"
)

func main() {
	fmt.Println(puppy.Barks())
	fmt.Println("Hello, exercise 14")
}
