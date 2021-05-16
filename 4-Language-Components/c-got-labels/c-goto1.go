package main 

import (
	"fmt"
)

func main(){
	// Local Variable Decleration
	var a int =1

	// Do Loop Executions
Loop:
	for a < 20{
		if a == 15{
			// Skip the Iteration
			a = a+1
			goto Loop

		}
		fmt.Printf("value of a: %d\n",a)
		a++
	}
}