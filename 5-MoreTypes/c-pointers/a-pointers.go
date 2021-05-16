package main

import (
	"fmt"
)

func main (){
	// Pointers are the refrences to defined variables
	// Some Operations like "Call by refrences can be done "

	// Every Variables is a Memory Location  and Every Location

	var  num1 int = 123 /* actual variable declaration */

	fmt.Println("\n vlaue" :%v",num1)
	fmt.Println("\n type" : %T",num1)
	fmt.Println("\n address location" :%x",num1)

	fmt.Println()

	var num1_ptr *int
	fmt.Printf("\n value" :%v",num1_ptr)
	fmt.Printf("\n type" :%T",num1_ptr)
	fmt.Printf("\n address location:%x,num_ptr)
	fmt.Println()

	num1_ptr =&num1
	fmt.Printf("\n value" :%v",num1_ptr)
	fmt.Printf("\n type" :%T",num1_ptr)
	fmt.Printf("\n address location:%x,num_ptr)
	fmt.Println()

	



	
}