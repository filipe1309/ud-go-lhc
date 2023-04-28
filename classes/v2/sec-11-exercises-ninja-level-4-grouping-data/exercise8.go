package main

import (
	"fmt"
)

func main() {
	xMap := map[string][]string{
		"bond_james":      {"Shaken not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}
	fmt.Println(xMap)

	for k, v := range xMap {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
