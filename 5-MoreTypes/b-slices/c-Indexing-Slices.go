package main

import (
	"fmt"
	"reflect"
)

func main() {

	primesArray := [6]int{2,3,5,7,11,13}
	
	fmt.Println(reflect.TypeOf(primesArray),reflect.TypeOf(primesArray).Kind(),primesArray)


	var slice1 []int = primesArray[1:4]
	fmt.Println(reflect.TypeOf(slice1),reflect.TypeOf(slice1).Kind(),slice1)

}