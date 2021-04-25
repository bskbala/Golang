package main

import (
	"fmt"
)

func main(){
	var number,remainder,temp int
	var reverse int =0

	fmt.Print("Enter ayn postive Integer:")
	fmt.Scan(&number)

	temp=number

	// for Loop Used Informat of While Loop
	for{
		remainder = number%10
		reverse =reverse*10 + remainder
		number/=10

		if(number == 0){
		break
		}
	}
	if(temp == reverse){
		fmt.Print("%d is a palindrome",temp)
	}else{
		fmt.Printf("%d is not a palindrome",temp)
	}



}
