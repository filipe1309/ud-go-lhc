package main

/**
 * In addition to the main goroutine, launch two additional goroutines.
 * Each additional goroutine should print something out.
 * Use waitgroups to make sure each goroutine finishes before your program exists.
 */

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Hello from the main goroutine")
	fmt.Println("Begin Number of CPUs:", runtime.NumCPU())
	fmt.Println("Begin Number of goroutines:", runtime.NumGoroutine())
	fmt.Println()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("G1:\tHello from the first goroutine")
		wg.Done()
	}()

	go func() {
		fmt.Println("G2:\tHello from the second goroutine")
		wg.Done()
	}()
	fmt.Println("Mid Number of goroutines:", runtime.NumGoroutine())
	fmt.Println()
	wg.Wait()

	fmt.Println("\nEnd Number of goroutines:", runtime.NumGoroutine())
	fmt.Println("About to exit")
}
