package main

/**
 * Starting with this code, sort the []user by
 * 	- age
 * 	- last
 * Also sort each []string "Sayings" for each user
 * 	- print everything in a way that is pleasant
 */

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("STRUCT SLICE:")
	fmt.Println(users)

	// your code goes here

	fmt.Println("Sort with sort.Slice:")

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	fmt.Println("BY AGE:")
	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})

	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, v2 := range v.Sayings {
			fmt.Println("\t", v2)
		}
	}

	fmt.Println("BY LAST:")
	sort.Slice(users, func(i, j int) bool {
		return users[i].Last < users[j].Last
	})

	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, v2 := range v.Sayings {
			fmt.Println("\t", v2)
		}
	}

	fmt.Println("-----------------------")
	fmt.Println("Sort with sort.Sort:")

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	fmt.Println("BY AGE:")
	sort.Sort(ByAge(users))

	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, v2 := range v.Sayings {
			fmt.Println("\t", v2)
		}
	}

	fmt.Println("BY LAST:")
	sort.Sort(ByLast(users))

	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, v2 := range v.Sayings {
			fmt.Println("\t", v2)
		}
	}
}
