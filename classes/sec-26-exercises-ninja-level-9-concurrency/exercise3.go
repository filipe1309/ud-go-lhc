package main

/**
 * Using goroutines, create an incrementer program
 * 	have a variable to hold the incrementer value
 * 	launch a bunch of goroutines
 * 		each goroutine should
 * 			read the incrementer value
 * 				store it in a new variable
 * 			yield the processor with runtime.Gosched()
 * 			increment the new variable
 * 			write the value in the new variable back to the incrementer variable
 * 	use waitgroups to wait for all of your goroutines to finish
 *	the above will create a race condition.
 *	Prove that it is a race condition by using the -race flag
 *	if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
 */

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func main() {
	incrementer := 0
	var wg sync.WaitGroup

	totalGoroutines := 100

	for i := 0; i < totalGoroutines; i++ {
		wg.Add(1)
		go func() {
			v := incrementer
			runtime.Gosched() // yield the processor
			v++
			incrementer = v
			fmt.Println("id:", goid(), "incrementer:", incrementer)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value:", incrementer)
}

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
