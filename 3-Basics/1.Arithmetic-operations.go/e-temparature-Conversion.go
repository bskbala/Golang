package main

import (
	"fmt"
)

func main() {
	const F float32 = 100
	const C float32 = (F - 32) * 5 / 9

	fmt.Println("F = ", F)
	fmt.Println("C = ", C)

	var F1 float32
	fmt.Print("Please Enter the Temperature in degrees Fahrenhit:")
	fmt.Scanf("%f", &F1)
	C1 := (F1 - 32) * 5 / 9
	fmt.Println("F1 = ", F1)
	fmt.Println("C1 =", C1)

}
