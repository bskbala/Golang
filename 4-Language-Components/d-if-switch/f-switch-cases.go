package main 

import  (
	"fmt"
)

func main () {
	fmt.Println("Enter Some number:")

	var value int
	fmt.Scanf("%d",&value)

	switch value % 2{
	case 0:
		fmt.Println("It is Even Number")
	default:
		fmt.Println("It is odd Number")
	}
}