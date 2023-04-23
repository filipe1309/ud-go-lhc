package main

/**
 * Starting with this code, marshal the []user to JSON.
 * There is a little bit of a curve ball in this hands-on exercise -
 * remember to ask yourself what you need to do to EXPORT a value from a package.
 */

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"first"`
	Age   int    `json:"age"`
}

func main() {
	u1 := user{"James", 32}
	u2 := user{"Moneypenny", 27}
	u3 := user{"M", 54}
	users := []user{u1, u2, u3}
	fmt.Println(users)

	// your code goes here
	// marshal = convert to JSON
	myJsonByteSlice, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(myJsonByteSlice))
}
