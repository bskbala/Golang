package main

import (
	"fmt"
)

// Note: If you want to compare two slices, then use range for loop to match each element or you can use DeepEqual function. 

// Multi-Dimensional Slice: Multi-dimensional 
// slice are just like the multidimensional array, except that slice does not contain the size. 
func  main () {
	// Creating Multi Dimensional Slices

	s1 :=[]int{
	{12,34}, 
	{56,47},
	{29,40},
	{46,78},
}

// Acessig Multi-dimensional slice
	fmt.Println("Slice 1:",s1)

// Creating Multi-Dimensional Slice
s1 := [][]int{{12, 34},
        {56, 47},
        {29, 40},
        {46, 78},
    }
 
    // Accessing multi-dimensional slice
    fmt.Println("Slice 1 : ", s1)
 
    // Creating multi-dimensional slice
    s2 := [][]string{
        []string{"Geeks", "for"},
        []string{"Geeks", "GFG"},
        []string{"gfg", "geek"},
    }
 
    // Accessing multi-dimensional slice
    fmt.Println("Slice 2 : ", s2)
 
}


}