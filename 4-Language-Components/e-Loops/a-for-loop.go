package main

import (
	"fmt"
)

func main(){
	// Method-1

	var i int
	for i =0; i<=10 ;i++ {
		fmt.Printf("%d\t",i)
	}
	fmt.Println()

	// Method -2
	for j:=0; j<=5 ; j++{
		fmt.Printf("%d\t",j)

	}
	fmt.Println()
	// Method -3
	j :=0
	for j<=5 {
		fmt.Printf("%d\t",j)
		j++
	}
	fmt.Println()
}
