package main

import (
	"fmt"
)


// defining the struct

type Car struct{
	Name,Model,Color string
	weighInkg float64
}
func main () {

	c :=Car{Name: "Ferrari",Model: "GTc4",
	Color: "Red",weighInkg: 1920,
	}

	// Accessing structs Fields
	// using the dot Operator
	fmt.Println("Car Name:",c.Name)
	fmt.Println("Car Color:",c.Color)

	// Assigning a new value
	// to a struct Field

	c.Color ="Black"

	// Displaying the result
	fmt.Println("Car:",c)
}