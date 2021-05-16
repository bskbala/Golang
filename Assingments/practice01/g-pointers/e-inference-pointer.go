package main

import (
	"fmt"
)

func main (){
	// Using var Keyward
	// we are not Definding
	// ay type with varaible

	var y =458

	// taking a pointer varaible using
	// we are not definding 
	// any type with varaibles

	var p = &y
	fmt.Println("value strored in y =" ,y)
	fmt.Println("Address of y =" ,&y)
	fmt.Println("Value Stored in Pointer varaible p=" ,p)
}