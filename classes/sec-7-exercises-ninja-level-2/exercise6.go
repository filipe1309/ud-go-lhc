package main

import (
	"fmt"
)

const (
	_                    = iota
	lastYear             = 2023 - iota
	lastLastYear         = 2023 - iota
	lastLastLastYear     = 2023 - iota
	lastLastLastLastYear = 2023 - iota
)

func main() {
	fmt.Println(lastYear)
	fmt.Println(lastLastYear)
	fmt.Println(lastLastLastYear)
	fmt.Println(lastLastLastLastYear)
}
