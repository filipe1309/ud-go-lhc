package main

/**
 * Starting with this code, unmarshal the JSON into a Go data structure.
 * Hint: https://mholt.github.io/json-to-go/
 */

import (
	libJson "encoding/json"
	"fmt"
)

// type secretAgents []struct {
// 	First   string   `json:"First"`
// 	Last    string   `json:"Last"`
// 	Age     int      `json:"Age"`
// 	Sayings []string `json:"Sayings"`
// }

type secretAgent struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	json := `[{"First":"James", "Last":"Bond", "Age":32, "Sayings":["Shaken, not stirred", "Any last wishes?", "Never say never"]}, {"First":"Miss", "Last":"Moneypenny", "Age":27, "Sayings":["James, it is soo good to see you", "Would you like me to take care of that for you, James?", "I would really prefer to be a secret agent myself."]},{"First":"M", "Last":"Hmmmm", "Age":54, "Sayings":["Oh, James. You didn't.", "Dear God, what has James done now?", "Can someone please tell me where James Bond is?"]}]`

	// your code goes here
	// var mySecretAgents secretAgents
	var mySecretAgents []secretAgent
	err := libJson.Unmarshal([]byte(json), &mySecretAgents)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mySecretAgents)
	fmt.Println(mySecretAgents[0].First)
}
