package main

import (
		"fmt"
		"reflect"
)

func main (){
	var val bool 
	fmt.Printf("val: %v\n",val)
	fmt.Printf("type : %T \n\n",val)

	var val2 = true
	fmt.Printf("val: %v\n",val2)
	fmt.Printf("type : %T \n\n",val2)

	val3 := 0 == 0
	fmt.Printf("val: %v\n",val3)
	fmt.Printf("type : %T \n\n",val3)

	const val4 = false
	fmt.Printf("val: %v\n",val4)
	fmt.Printf("type : %T \n\n",val4)
	fmt.Println("reflect.TypeOf(val4) =",reflect.TypeOf(val4))


}