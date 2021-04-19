package main

import "fmt"

// Package -Level Declaration
var name = "pacakge-level"

func main() {
	// Function-Level Declaration
	fmt.Println("In Main() : name =", name)

	var name = "function-level"
	fmt.Println("In main () : name", name)

}

func myfunc() {
	fmt.Println("In main () : name =", name)
}

// Note: Function -level declartions will  be prioritized,
// comapared to Pcakege Level Decalartion
