package main

import (
	"fmt"
)

func main () {
	// taking variable
	// of integer type

	var V int = 100

	// taking a pointer
	// of the Integer type

	var pt1 * int =& V

	// taking pointer to
	// pointer to pt1
	// storing the address
	// of pt1 into pt2

	var pt2 ** int = &pt1

	fmt.Println("The value of varaiable V is =",V)
	fmt.Println("Address of Variable V is =",V)
	fmt.Println("The Value of pt1 is =",pt1)
	fmt.Println("Address of pt1 is =", &pt1)

	fmt.Println("The Value of pt2 is =",pt2)

	// Dereferrencing the 
	// Pointer to Pointer

	fmt.Println("Value at the address of pt2 is or *pt2 =", *pt2)

	// Double Pointer will give the value of varaibel V
	fmt.Println("*(Value at  the address of pt2 is )or ** pt2 =",**pt2)


}