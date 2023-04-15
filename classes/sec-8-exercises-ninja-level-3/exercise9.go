package main

import (
	"fmt"
)

func main() {
	favSport := "running"
	switch favSport {
	case "rugby":
		fmt.Println("I like rugby")
	case "football":
		fmt.Println("I like football")
	case "tennis":
		fmt.Println("I like tennis")
	case "running":
		fmt.Println("I like running")
	}
}
