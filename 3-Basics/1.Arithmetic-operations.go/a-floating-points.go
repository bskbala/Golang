package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34
	fmt.Println(Avogadro, reflect.TypeOf(Avogadro))
	fmt.Println(Planck, reflect.TypeOf(Planck))
	fmt.Println()

	postiveInfinity := math.Inf(0)
	fmt.Println(postiveInfinity, reflect.TypeOf(postiveInfinity))
	negativeInfinity := math.Inf(-98)
	fmt.Println(negativeInfinity, reflect.TypeOf(negativeInfinity))
	garbagevalue := math.Inf
	fmt.Println(garbagevalue, reflect.TypeOf(garbagevalue))
	fmt.Println()

	x := 1
	fmt.Println("math.Exp(float64(x)):",math.Exp(float64(x)))
	x = 2
	fmt.Println("math.Exp(float64(x)):",math.Exp(float64(x)))
	x = 3
	fmt.Println("math.Exp(float64(x)):",math.Exp(float64(x)))


}
