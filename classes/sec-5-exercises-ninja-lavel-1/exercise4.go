package main

import (
	"fmt"
)

type myOwnType int // underlying type is int
var x myOwnType

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
