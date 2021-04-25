package main

import (
	"fmt"
	"sort"
)

func main (){

	slc1 :=[]string{"python","java","C#","Go","Ruby"}
	slc2 :=[]int{45,65,23,90,33,21,56,78,89}

	fmt.Println("Before Sorting")
	fmt.Println("Slice 1:",slc1)
	fmt.Println("Slice 2:",slc2)

	// Performing sort operation on the slice using sort function

	sort.Strings(slc1)
	sort.Ints(slc2)
	
	fmt.Println("\n After Sorting")
	fmt.Println("Slice 1:",slc1)
	fmt.Println("Slice",slc2)



}