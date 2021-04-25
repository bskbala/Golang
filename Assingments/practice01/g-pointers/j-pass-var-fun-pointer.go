package  main

import (
	"fmt"
)

// taking a function with integer
// type pointer as a parameter

func ptf(a *int){
	// dereferencing
	*a =748
}

func main() {

	// taking Normal variable
	var x =100
	fmt.Println("the value of x before function call is :%d\n",x)

	// calling the function by 
	// passing the address of
	// the variable x

	ptf(&x)
	fmt.Println("the  valueof x after function call  is :%d\n",x)


}