package main

/**
 *	- Create the following variables with the following scopes:
 *		- package level
 *			- Create outside of func main
 *			- Use the:
 *				- var keyword
 *				- const keyword
 *		- block level
 *			- Create inside of func main
 *			- Use the short declaration operator
 *	- Use the variables in func main
 */

import "fmt"

var myPackageVariable string = "This is a package level variable"

const myPackageConstant string = "This is a package level constant"

func main() {
	myLocalVariable := "This is a local variable"
	fmt.Println("myPackageVariable:", myPackageVariable)
	fmt.Println("myPackageConstant:", myPackageConstant)
	fmt.Println("myLocalVariable:", myLocalVariable)
}
