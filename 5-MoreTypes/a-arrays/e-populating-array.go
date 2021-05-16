package main

import (
	"fmt"
)

func main() {
	var numbers [10] int

	// Getting values  in runtime

	fmt.Println("Enter 10 numbers space seprated:")
	fmt.Scanf("%d", &numbers)
	fmt.Println("len(numbers)=",len(numbers))

	for index :=0; index <len(numbers);index++{
		fmt.Scanln( &numbers[index])
		fmt.Println(index,len(numbers),numbers[index],numbers)

	}
	fmt.Println("numbers=",numbers)


}