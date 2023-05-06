package main

/**
 * Create a program that uses iota to calculate the size of each measurement of bytes
 * 	- https://golang.org/ref/spec#Iota
 * 	- https://golang.org/doc/effective_go#constants
 * 	- KB, MB, GB, TB, PB, EB
 * - Show the size of each in DECIMAL and BINARY using fmt.Printf
 * - hint: use int and not float64 as shown in effective_go#constants (we aren't going up to zettabytes)
 * - hint: the starting code for this is already in effective_go#constants
 */

import "fmt"

type ByteSize int64

const (
	_           = iota             // 0
	KB ByteSize = 1 << (10 * iota) // 1 << (10 * 1) = 1 << 10 = 1024
	MB                             // 1 << (10 * 2) = 1 << 20 = 1048576
	GB                             // 1 << (10 * 3) = 1 << 30 = 1073741824
	TB                             // 1 << (10 * 4) = 1 << 40 = 1099511627776
	PB                             // 1 << (10 * 5) = 1 << 50 = 1125899906842624
	EB                             // 1 << (10 * 6) = 1 << 60 = 1152921504606846976
)

func main() {
	fmt.Println(KB, MB, GB, TB, PB, EB)
}
