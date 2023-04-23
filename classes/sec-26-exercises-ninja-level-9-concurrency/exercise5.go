package main

/**
 * Fix the race condition you created in the previous exercise by using package atomic
 */

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func main() {
	var incrementer int64
	var wg sync.WaitGroup

	totalGoroutines := 100

	for i := 0; i < totalGoroutines; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println("id:", goid(), "incrementer:", atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value:", incrementer)
}

func goid() int {
	var buf [64]byte                                                                // 64 bytes should be enough for most goroutine ids
	n := runtime.Stack(buf[:], false)                                               // runtime.Stack() returns the number of bytes written to buf
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0] // get the goroutine id from the stack trace
	id, err := strconv.Atoi(idField)                                                // convert the goroutine id from string to int
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
