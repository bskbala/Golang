package main 

import (
	"fmt"
)

func main () {
	fmt.Println("Enter  Some number")
	var value int32
	fmt.Scanf("%d",&value)

	switch value{
	case 1: 
		fmt.Println("==========ONE======")
	case 2 :
		fmt.Println("==========TWO======")
	case 3:
		fmt.Println("==========THREE======")
	default:
		fmt.Println("==INVALID NUMBER")

	}

}