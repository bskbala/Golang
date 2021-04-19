package main

import (
	"fmt"
)

// Display the largest among the Two numbers
func main (){
	var num1 ,num2 int
	fmt.Println("Please Enter  Number1:")
	fmt.Scanf("%d",&num1)
	fmt.Println("Please Enter  Number2:")
	fmt.Scanf("%d",&num2)

	fmt.Println("num1=",num1)
	fmt.Println("num2=",num2)

	if  num1 >num2 {
		fmt.Println("Num1 is greater ")
	}else if num1 < num2 {
		fmt.Println("Num2 is greater")
	}else{
		fmt.Println("num2 is grater")
	}






}