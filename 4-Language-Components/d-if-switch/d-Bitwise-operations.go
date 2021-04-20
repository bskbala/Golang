package main

import (
	"fmt"
)

func main() {
	num1 := 20
	num2 := 12

	fmt.Println("num1 =%v\n", num1)
	fmt.Println("num1 =%b\n\n", num1)

	fmt.Println("num2 =%v\n", num2)
	fmt.Println("num2 =%b\n\n", num2)

	fmt.Println("num1 & num 2 ==", num1&num2)
	fmt.Println("num1 | num 2 ==", num1|num2)
	fmt.Println("num1 ^ num 2 ==", num1^num2)

}
