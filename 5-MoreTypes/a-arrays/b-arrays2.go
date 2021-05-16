package main

import (
	"fmt"
	"reflect"
)

func main (){
	a1 := [3]int{11,22,23}
	fmt.Println("%v- %[1]T-%v\n",a1,reflect.TypeOf(a1).Kind())

	a2 := [...]int{11,22,23}
	fmt.Println("%v -%[1]T-%v\n",a1,reflect.TypeOf(a2).Kind())

	a3 := []int{11,22,23}
	fmt.Println("%v- %[1]T-%v\n",a1,reflect.TypeOf(a3).Kind())

	fmt.Println("a1 == a2:",a1 == a2 )
}