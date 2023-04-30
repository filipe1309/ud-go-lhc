package main

/**
 * We are going to learn about testing next.
 * For this exercise, take a moment and see how much you can figure out
 * about testing by reading the testing documentation (https://golang.org/pkg/testing/)
 * & also Caleb Doxsey’s article (https://www.golang-book.com/books/intro/12) on testing.
 */

import "testing"

func TestSqrt(t *testing.T) {
	var v float64
	v, _ = Sqrt(1024)
	if v != 42 {
		t.Error("Expected 1.5, got ", v)
	}
}
