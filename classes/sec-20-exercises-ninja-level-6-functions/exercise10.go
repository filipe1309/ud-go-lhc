package main

/**
 * Closure is when we have “enclosed” the scope of a variable in some code block.
 * For this hands-on exercise, create a func which “encloses” the scope of a variable:
 */

import (
	"fmt"
)

func main() {
	myFuncThatIncrementEnclosuredVar := myIncrementFunc()
	fmt.Println(myFuncThatIncrementEnclosuredVar())
	fmt.Println(myFuncThatIncrementEnclosuredVar())
	fmt.Println(myFuncThatIncrementEnclosuredVar())
	fmt.Println(myFuncThatIncrementEnclosuredVar())

	mySecondFuncThatIncrementEnclosuredVar := myIncrementFunc()
	fmt.Println(mySecondFuncThatIncrementEnclosuredVar())
	fmt.Println(mySecondFuncThatIncrementEnclosuredVar())
	fmt.Println(mySecondFuncThatIncrementEnclosuredVar())
	fmt.Println(mySecondFuncThatIncrementEnclosuredVar())
	fmt.Println(mySecondFuncThatIncrementEnclosuredVar())
	fmt.Println(mySecondFuncThatIncrementEnclosuredVar())
	fmt.Println(mySecondFuncThatIncrementEnclosuredVar())
}

func myIncrementFunc() func() int {
	myEnclosuredVar := 0
	return func() int {
		myEnclosuredVar++
		return myEnclosuredVar
	}
}
