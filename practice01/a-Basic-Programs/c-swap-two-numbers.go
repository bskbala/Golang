package main

import "fmt"

func main (){
	var first int
	fmt.Printf("Enter the First  Number ")
	fmt.Scanf("%d",&first)
	var second int
	fmt.Printf("Enter the second  Number ")
	fmt.Scanf("%d",&second)
	first =first -second
	second =first+second
	first =first+second
	fmt.Printf("first Number:" ,first)
	fmt.Printf("second number :",second)

	
}