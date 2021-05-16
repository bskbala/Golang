package main

import (
	"fmt"
	"reflect"
)

func main (){

	array1 := [3]int{2,3,5}
	fmt.Println(reflect.TypeOf(array1).Kind(),array1)

	array2 :=[...]bool{false,true,true}
	fmt.Println(reflect.TypeOf(array2).Kind(),array2)

	var array3 = []bool{false,true,true}
	fmt.Println(reflect.TypeOf(array3).Kind(),array3)

	mySlice1 := []bool{false,true,true}
	fmt.Println(reflect.TypeOf(mySlice1).Kind(),mySlice1)


	mySlice2 := make([]string,3)
	fmt.Println(reflect.TypeOf(mySlice2).Kind(),mySlice2)



}