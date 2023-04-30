package main

/**
 * Starting with this code(http://play.golang.org/p/wlEM1tgfQD),
 * use the sqrtError struct as a value of type error.
 * If you would like, use these numbers for your
 * lat "50.2289 N"
 * long "99.4656 W"
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
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code hereß
		// e := errors.New("Math error: square root of negative number")
		e := fmt.Errorf("Math error: square root of negative number: %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}
