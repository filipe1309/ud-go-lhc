package main

import (
	"fmt"
)

func main() {
	myAnonymousStruct := struct {
		something    string
		anotherThing int
		someMap      map[string]int
		someSlice    []string
	}{
		something:    "something",
		anotherThing: 2,
		someMap: map[string]int{
			"one": 1,
			"two": 2,
		},
		someSlice: []string{"one", "two", "three"},
	}

	fmt.Println(myAnonymousStruct)
	fmt.Println(myAnonymousStruct.something)
	fmt.Println(myAnonymousStruct.anotherThing)
	fmt.Println(myAnonymousStruct.someMap)
	fmt.Println(myAnonymousStruct.someSlice)

	for k, v := range myAnonymousStruct.someMap {
		fmt.Println(k, v)
	}

	for i, v := range myAnonymousStruct.someSlice {
		fmt.Println(i, v)
	}
}
