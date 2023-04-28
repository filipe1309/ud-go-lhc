package main

/**
 * Create a program that prints out your OS and ARCH.
 * Use the following commands to run it
 * 	go run
 * 	go build
 * 	go install
 */

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
}
