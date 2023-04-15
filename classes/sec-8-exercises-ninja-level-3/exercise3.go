package main

import (
	"fmt"
	"time"
)

func main() {
	actualYear := time.Now().Year()
	birthYear := 1989
	for birthYear <= actualYear {
		fmt.Println(birthYear)
		birthYear++
	}
}
