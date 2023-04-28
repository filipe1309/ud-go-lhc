package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")
	foo()
	fmt.Println("Hello, again.")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited.")
}
