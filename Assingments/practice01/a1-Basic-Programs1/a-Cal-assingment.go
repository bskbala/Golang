package main

import "fmt"

func main() {
	fmt.Print("Please enter first number: ")
	var n1 int
	fmt.Scanln(&n1) // take input from user
	fmt.Print("Please enter Second number: ")
	var n2 int
	fmt.Scanln(&n2) // take input from user

	// Addition
	result := n1 + n2
	fmt.Println("Sum of two numbers: ", result)

	// subtraction
	result = n1 - n2
	fmt.Println("Subtraction of two numbers: ", result)

	// Multiplication
	result = n1 * n2
	fmt.Println("Multiplication of two numbers: ", result)

	// Division
	result = n1 % n2

	fmt.Println("Division of two numbers: ", result)

}
