package main

import "fmt"

func main() {

	// variable declaration  & Initialization

	// Method -1
	var x string
	x = "hello world"
	fmt.Println("x=", x)

	// Method-2
	var y string = "Hello world"
	fmt.Println("y=", y)

	// Method -3 type is not Necessary beacuse  the Go Complier is able to
	// Infer the type based on the Literal value assinged  to the varaible.
	z := "hello world"

	// Type Inference - variable's type is Inferred form the value on the right hand Side

	fmt.Println("z=", z)
	i := 42
	f := 3.142
	g := 0.867 + 0.51
	fmt.Println(i, f, g)

	// Bulk Variable Declartion  Cum Initialization

	var (
		a = 5
		b = 9
		c = 15
	)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
}
