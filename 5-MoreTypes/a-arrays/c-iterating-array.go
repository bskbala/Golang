package main

import ( 
	"fmt"
)

func main(){
	oddNumbers := [...]int{1,3,5,7,9,11,13}

	// Indexing arrays

	fmt.Println("Fifth Element in array is ", oddNumbers[5])
	fmt.Println("Fifth Element in array is ", oddNumbers[6])

	for index := 0; index <len(oddNumbers); index++{
		fmt.Printf("In array %v , Value at Position %d is %2d\n",oddNumbers, index,oddNumbers[index])
	}
	fmt.Println()

	// Iterataing Using Range

	for index,value := range oddNumbers{
		fmt.Printf("In array %v , Value at Position %d is %2d\n",oddNumbers, index,value)
	}


}