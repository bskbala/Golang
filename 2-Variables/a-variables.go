package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println("a=", a)

	var b = 12312323
	fmt.Println("b=", b)

	var c float64 = 123.45
	fmt.Println("c= ", c)

	// More than One Variable can be Declared

	var f, g int = 30, 40
	fmt.Println("f =", f, "g =", g)

	// Go will infer the type of Initialzed Variable

	var h = true
	fmt.Println("h=", h)

	// Variables Declared  Without Initilazation
	// Zero Value for the Int -> 0
	// Zero value for th str -> ""

	var i int
	fmt.Println("i=", i)

	var (
		j float32
		k float64
	)
	fmt.Println("j:", j)
	fmt.Println("k:", k)

}
