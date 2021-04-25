package main

import (
	"fmt"
)

type Employee struct{
	firstName string
	lastName string
	age int
	salary int
}
func  main () {
	// Creating struct specifying field names
	emp1 := Employee{
		firstName: "Sam",
		age :25,
		salary: 500,
		lastName:"kumar",
	}

	// Creating struct with out specifying Fields names

	emp2 := Employee{"randura","apparao",25,599}

	fmt.Println("Employee 1",emp1)
	fmt.Println("Employee 2",emp2)
}