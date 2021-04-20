package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter Some number (0-9) :")
	var value int
	fmt.Scanf("%d", &value)

	switch value{
	case 0:
		fmt.Println("it is zero")
	case 1,3,5,7,9:
		fmt.Println("It is odd number")
	case 2,4,6,8:
		fmt.Println("it is Even  Number")
	default:
		fmt.Println("Ther is  no number ")
	}
}