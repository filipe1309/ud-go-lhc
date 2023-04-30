package main

/**
 * Create a struct “customErr” which implements the builtin.error interface.
 * Create a func “foo” that has a value of type error as a parameter.
 * Create a value of type “customErr” and pass it into “foo”.
 */

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.info)
}

func foo(e error) {
	fmt.Printf("foo - %T\n", e)
	fmt.Println("foo -", e)
	// fmt.Println("foo -", e.info)             // this doesn't work
	fmt.Println("foo -", e.(customErr).info) // get info from customErr struct with type assertion
}

func main() {
	myCutomError := customErr{info: "This is my custom error"}
	foo(myCutomError)
}
