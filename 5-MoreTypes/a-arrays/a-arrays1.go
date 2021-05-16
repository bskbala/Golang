package main

import (
	"fmt"
)

/*
Array - datastructure to store a collection  of data of the same type
Declaration
 var array_name[lenght]Type
 or
 var array_name[Length]Type{item1,item2,item3,.....item.N}
 
 short Decleration

 array_name :=[length]Type {Item1,item2,item3...Item.nn}

*/
func main(){
	//Declearation Off arrays

	var emptyArray1 [3] int
	fmt.Println("emptyArray1 = ",emptyArray1)

	var emptyArray2 [3] string
	fmt.Println("emptyArray2 = ",emptyArray2)

	var emptyArray3 [3] bool
	fmt.Println("emptyArray3 = ",emptyArray3)

	fmt.Println()

// Decleration and Initialzation

	var myArray1 [3]int =[3]int{11,22,33}
	fmt.Println("myArray1 =", myArray1)

	var myArray2  =[3]int{11,22,33}
	fmt.Println("myArray2 =", myArray2)

	 myArray3 := [3]int{11,22,33}
	fmt.Println("myArray3 =", myArray3)
	fmt.Println()


// Multi-line Initialization

	greetings1 :=[4]string{"Good Morning","Good AfterNoon","Good Evening"}
	fmt.Println("greetings1 = ", greetings1)

	greetings2 :=[4]string{
	"Good Morning",
	"Good AfterNoon",
	"Good Evening",
}
	fmt.Println("greetings2 = ", greetings2)
	fmt.Println()

	greetings3 :=[...]string{
		"Good Morning",
		"Good AfterNoon",
		"Good Evening",
	}
		fmt.Println("greetings3 = ", greetings3)
		fmt.Println("len(greetings1) = ", len(greetings1))
		fmt.Println("len(greetings2) = ", len(greetings2))
		fmt.Println("len(greetings3) = ", len(greetings3))
		fmt.Println()



		



}