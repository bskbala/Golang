package main

import "fmt"

/*
Purpose of Enum:
    1.Grouping and Expecting  only some related values
	2.Sharing  common Behavior
	3.Aviods using invalid Values
	4.To Increase the code readability  and the maintainability
*/

func main() {
	// const (
	// 	a1 = 10
	// 	a2 = 11
	// 	a3 = 12
	// 	a4 = 13
	// )

	/*
	 Go does not have theEnumarate Types

	 Instead,you can use the specail name "Iota"

	*/
	const (
		C0 = iota
		C1 = iota
		C2 = iota
	)

	fmt.Println(C0, C1, C2, C2)
	fmt.Println()

	// ===========================
	// When an initialization Expression is Omitted for
	// it resuses the Preceding Expression

	const (
		red = iota
		blue
		green
	)
	fmt.Println("red =", red)
	fmt.Println("blue =", blue)
	fmt.Println("green =", green)
	fmt.Println()

	// Counting Backward

	const (
		max = 10
	)
	const (
		a = (max - iota)
		b
		c
	)

	fmt.Println(a, b, c)
}
