package main

/**
 * Write a program that uses the following:
 * 	- var for zero value
 * 	- short declaration operator
 * 	- multiple initializations
 * 	- var when specificity is required
 * 	- blank identifier
 * print items as necessary to make the program interesting
 */

import "fmt"

func main() {
	fmt.Println("var for zero value")
	var a int
	fmt.Println(a)

	fmt.Println("short declaration operator")
	b := 42
	fmt.Println(b)

	fmt.Println("multiple initializations")
	var c, d int = 1, 2
	fmt.Println(c, d)

	fmt.Println("var when specificity (the type) is required")
	var e int = 3
	fmt.Println(e)

	fmt.Println("blank identifier")
	_, f := 4, 5
	fmt.Println(f)
}
