package main 

import (
	"fmt"
)

func main () {
	// taking a normal variable 
	var x int =5748

	// Decleration of Pointer

	var p* int

	// Initializzation  of Pointer

	p=&x

	// Displaying the result

	fmt.Println("Value stored in x =",x)
	fmt.Println("Addresss of x =",&x)
	fmt.Println("Value Stored in variables p=",p)

}