package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	z := x << 1 // shift bits to the left by 1
	fmt.Printf("%d\t%b\t%#x\n", z, z, z)
}
