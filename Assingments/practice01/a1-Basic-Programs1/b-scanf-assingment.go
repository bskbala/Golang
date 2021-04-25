package main

import "fmt"

func main() {

	
	
	fmt.Println("please enter first  Number:")
	var num1 int
	fmt.Scanf("%d", &num1)
	fmt.Println()
	fmt.Println("Please enter  second Number:")
	var num2 int
	fmt.Scanf("%d", &num2)

	result := num1 + num2
	fmt.Println("sum of two Numbers =", result)

}
