package main

import (
	"fmt"
)

func main() {

	var y =458
	var p =&y

	fmt.Println("Value stored in Y before changing =",y)
	fmt.Println("Address of y =",&y)
	fmt.Println("Value store in pointer varaible p=",p)

	// This is derefrencing a pointer
	// using * operator before a pointer
	// varaible to access the value stored
	// at the varaible at which it is Pointing 

	fmt.Println("Value stored in y(*p) Before Chaning =",*p)
	// Changing the value of Y by assgning
	// the new value  to the Pointer

	*p = 500
	fmt.Println("Value stored in y(*p)after changing=",y)


}