package main

import (
	"fmt"
	"strings"
)

func main() {

	// String Comparsion

	fmt.Println("\n====String Comparsion ====")
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
	fmt.Println()

	fmt.Println(strings.Compare("apple", "apple"))
	fmt.Println(strings.Compare("apple", "apdama"))
	fmt.Println()
}
