package main

import (
	"fmt"
)
/*
  when array elements of an array are arrays,thrn it is called multi-dimensional arrays
*/
func main(){

	//   Multi-dimensions arrays

	 mulArr := [3][2]int{
		[2]int{1,2},
		[2]int{3,4},
		[2]int{5,6},
	}
	fmt.Println("mulArr=" ,mulArr)

	mulArr2 := [...][2]int{
		[2]int{1,2},
		[2]int{3,4},

	}
	fmt.Println("mulArr2=" ,mulArr2)

	// Indexing Multi-Dimensional arrays

	for index,subArray := range mulArr2{
		fmt.Println("In array",index,subArray)
		for subIndex,element := range subArray{
			fmt.Println(subIndex,element)
			fmt.Printf("mulArr2[%d][%d]=%d\n",index,subIndex,element)
		}
	}
}