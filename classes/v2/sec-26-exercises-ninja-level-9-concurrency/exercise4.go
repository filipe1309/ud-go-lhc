package main

/**
 * Fix the race condition you created in the previous exercise by using a mutex
 * it makes sense to remove runtime.Gosched()
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
	var mu sync.Mutex

	totalGoroutines := 100

	for i := 0; i < totalGoroutines; i++ {
		wg.Add(1)
		go func() {
			// Mutex
			mu.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println("id:", goid(), "incrementer:", incrementer)
			mu.Unlock()
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
