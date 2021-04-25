package main

import (
	"fmt"
)

func main (){
	var slice1 =[]string{"india","canada","usa"}
	var slice2 =[]string{"africa","austrila"}

	slice2 =append(slice2,slice1...)
	fmt.Println("slice1:",slice1)
	fmt.Println("slice2:",slice2)
}