package main

/**
 * We are going to learn about testing next.
 * For this exercise, take a moment and see how much you can figure out
 * about testing by reading the testing documentation (https://golang.org/pkg/testing/)
 * & also Caleb Doxsey’s article (https://www.golang-book.com/books/intro/12) on testing.
 */

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("Math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := Sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code hereß
		// e := errors.New("Math error: square root of negative number")
		e := fmt.Errorf("Math error: square root of negative number: %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}
