package main

import (
	"fmt"
)

// Defining a struct  Type

type Address struct{
	Name string
	city string
	Pincode int

}

func main() {
	// Declaring a varaible of a "struct" type
	// All the  struct Fields are initialzed
	// with  their Zero value

	var  a Address
	fmt.Println(a)

	// Declaring and Initalizaing a
	// struct using  a struct Literial
	a1 :=Address{"Akshay","Dehradun",3623572}
	fmt.Println("Address1:",a1)

	// Naming Fields While
	// Initalizing a struct

	a2 := Address{Name: "Anikaa", 
			city: "Ballia",
			Pincode: 277001}
  
    fmt.Println("Address2: ", a2)
  
    // Uninitialized fields are set to
    // their corresponding zero-value
    a3 := Address{Name: "Delhi"}
    fmt.Println("Address3: ", a3)

}