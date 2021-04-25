/*
What is a struct?
A struct is a user-defined type that represents a collection of fields. 
It can be used in places where it makes sense to group the data into a single unit rather than having each of them as separate values.

*/

package main

import (
	"fmt"
)


type Employee struct{
	firstName,lastName string 
	age ,salary int
}

// Main Function
func main(){

	// Taking Pointer to Struct

	emp8 :=&Employee{"Sam","Anderson",55,6000}

	// emp8.firstName is used to access
	// the field firstName
	fmt.Println("FirstName:",emp8.firstName)
	fmt.Println("lastname:",emp8.lastName)
	fmt.Println("Age:",emp8.age)
	fmt.Println("salary:",emp8.salary)


}