package main

/**
 * Using the code you wrote in the previous hands-on exercise:
 *	- Look in the $GOPATH/bin
 *		- Launch another terminal
 *		- See the "go path" environment variable - either of these
 *			- echo $GOPATH
 *			- go env GOPATH
 *		- Navigate to the $GOPATH/bin folder
 *			- cd $GOPATH/bin
 *			- ls -la
 *	- "go install" your program (on the terminal)
 *		- Look at the executable in the $GOPATH/bin folder
 *		- ls -la
 *		- go install
 *	- Run the executable in the $GOPATH/bin folder
 *		- ./<name-of-executable>
 *	- Remove the executable in the $GOPATH/bin folder
 *		- rm <name-of-executable>
 *		- ls -la
 *		- If you accidentally delete everything, you will need to reinstall your tooling in VS Code
 *		- If you messed it all up, reinstall Go
 */

import "fmt"

func main() {
	fmt.Println("Hello, exercise 12")
}
