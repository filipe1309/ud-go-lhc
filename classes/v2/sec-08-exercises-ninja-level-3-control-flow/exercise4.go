package main

import (
	"fmt"
	"time"
)

func main() {
	actualYear := time.Now().Year()
	birthYear := 1989
	for {
		if birthYear > actualYear {
			break
		}
		fmt.Println(birthYear)
		birthYear++
	}
}
