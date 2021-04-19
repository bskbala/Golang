package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Chantaing value after Declaration

	var x string
	x = "first"
	fmt.Println("x=", x)

	x = "second"
	fmt.Println("x=", x)

	x = x + "third"
	fmt.Println("x=", x)

	// Compound Operation

	x += "Fourth"
	fmt.Println("x=", x)

	// Introducing the  Reflect  package

	fmt.Println("x               =", x)
	fmt.Println("reflect.ValueOf(x)  =", reflect.ValueOf(x))
	fmt.Println("reflect.TypeOf(x)  =", reflect.TypeOf(x))

}
