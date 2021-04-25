package main

import (
	"fmt"
	
)

func main() {
	// creating a map
	// Using make() function

	var my_map =make(map[float64]string)
	fmt.Println(my_map)

	// As w already know that make() function
	// always returns a map which is initialized
	// So,wee can add values in it 
	my_map[1.3] ="saikumar"
	my_map[1.5] ="Dollar"
	fmt.Println(my_map)
}