package main

/**
 * Using the code you wrote in the previous exercise:
 *	- Build your program for windows
 *		- GOOS=windows GOARCH=amd64 go build
 *	- Build your program for mac
 *		- GOOS=darwin GOARCH=arm64 go build
 *	- Build your program for linux
 *		- GOOS=linux GOARCH=arm64 go build -o <name-of-executable>
 */

import "fmt"

func main() {
	fmt.Println("Hello, exercise 12")
}
