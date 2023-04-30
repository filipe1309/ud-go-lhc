package main

/**
 * Starting with this code (https://play.golang.org/p/9a1IAWy5E6).
 * Create a custom error message using “fmt.Errorf”.
 */

import (
	"encoding/json"
	// "errors"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJson(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}

func toJson(a any) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		// return []byte{}, fmt.Errorf("JSON did not marshal: %v", err)
		// return []byte{}, errors.New(fmt.Sprintf("JSON did not marshal: %s", err))
		return []byte{}, fmt.Errorf("JSON did not marshal: %w", err)
	}
	return bs, nil
}
