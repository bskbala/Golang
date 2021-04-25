package main  

import (
	"fmt"
)

func main() {
	names := [4]string{

		"saikumar",
		"balam",
		"bhargavi",
		"srilekha",
		
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a,b)


	// slices are Mutable
	b[0] ="XXX"
	fmt.Println(a,b)
	fmt.Println(names)
}