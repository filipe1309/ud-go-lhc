package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d is divisible by 3 and 5\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d is divisible by 3\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d is divisible by 5\n", i)
		} else {
			fmt.Printf("%d is not divisible by 3 or 5\n", i)
		}
	}
}
