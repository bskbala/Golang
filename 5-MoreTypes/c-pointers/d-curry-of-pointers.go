package main


import (

	"fmt"
)
func main(){

	var intVar int
	var pointerVar * int
	var pointerToPointerVar **int

	fmt.Println("\nInital Defaults:")
	fmt.Println("intVar   :",intVar)
	fmt.Println("pointerVar :", pointerVar)
	fmt.Println("pointerToPointerVar:",pointerToPointerVar)

	intVar =100
	pointerVar =&intVar
	pointerToPointerVar =&pointerVar


	fmt.Println("\nAfter Assingments:")
	fmt.Println("intVar   :",intVar)
	fmt.Println("pointerVar :", pointerVar)
	fmt.Println("pointerToPointerVar:",pointerToPointerVar)


}