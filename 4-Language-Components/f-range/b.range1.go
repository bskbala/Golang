package main

import (
	"fmt"
)

func main(){

	// value := nil
	// for index,val := range value{
	// 	fmt.Println(index,val) 
	// }

	// Looping Over an Array
	nums :=[]int{11,22,33,44,55,66,77}
	for index,val := range  nums{
		fmt.Println(index,val)
	}

	// Looping over a map -both key and value

	kvs :=map[string]string{"a":"apple","b":"banana","c":"cat"}

	for range kvs{
		fmt.Println("-in loop")
	}

	// Looping Over a Map -Only Key

	for k:= range kvs{

		fmt.Println("keys:", k)
	}

}