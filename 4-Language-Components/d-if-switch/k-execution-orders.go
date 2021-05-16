package main

import (
	"fmt"
)

//Foo Prints and Returns n
func Foo(n int) int {
	fmt.Println("n=",n)
	return n
}
func main (){
	// Foo(1)
	// Foo(2)
	// Foo(3)
	// Foo(4)

	switch Foo(2){
	case Foo(1),Foo(2),Foo(3):
		fmt.Println("First Case")
	case Foo(4):
		fmt.Println("Second Case")
	default:
		fmt.Println("This is Nibba")
	}
}