package main

import (
	"fmt"
)

func main() {
	equal := 42 == 42
	less_than_or_equal := 42 <= 42
	greater_than_or_equal := 42 >= 42
	not_equal := 42 != 42
	less_than := 42 < 42
	greater_than := 42 > 42
	fmt.Println("equal:", equal)
	fmt.Println("less_than_or_equal:", less_than_or_equal)
	fmt.Println("greater_than_or_equal:", greater_than_or_equal)
	fmt.Println("not_equal:", not_equal)
	fmt.Println("less_than:", less_than)
	fmt.Println("greater_than:", greater_than)
}
