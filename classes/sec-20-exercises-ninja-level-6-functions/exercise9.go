package main

/**
 * A callback is when we pass a func into a func as an argument. For this exercise,
 * pass a func into a func as an argument
 */

import (
	"fmt"
)

func main() {
	myFunc(myOtherFunc)
}

func myFunc(myCallbackFunc func(xi []int) int) {
	fmt.Println("My function")
	fmt.Println(myCallbackFunc([]int{1, 2, 3, 4, 5}))
}

func myOtherFunc(xi []int) int {
	fmt.Println("My callback function")
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
