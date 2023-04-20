package main

/**
 * Create a type SQUARE
 * Create a type CIRCLE
 * attach a method to each that calculates AREA and returns it
 * circle area= π r 2
 * square area = L * W
 * create a type SHAPE that defines an interface as anything that has the AREA method
 * create a func INFO which takes type shape and then prints the area
 * create a value of type square
 * create a value of type circle
 * use func info to print the area of square
 * use func info to print the area of circle
 */

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("Area: %v of %T\n", s.area(), s)
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	mySquare := square{
		side: 10,
	}

	myCircle := circle{
		radius: 10,
	}

	info(mySquare)
	info(myCircle)
}
