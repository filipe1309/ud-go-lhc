package main

import (
	"fmt"
)

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for i, v := range x {
		fmt.Println(i)
		for _, v2 := range v {
			fmt.Printf("\t%v", v2)
		}
		fmt.Println()
	}
}
