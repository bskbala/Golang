package main

import (
	"fmt"
)

func main(){
	var val int =42
	var valPtr * int = &val

	fmt.Println("val=", val,"valPtr=",*&valPtr)
	fmt.Println("valPtr=",valPtr,"&val=",&val)

	// Updating Value Directly
	val =66
	fmt.Println("val=", val,"valPtr=",*&valPtr)
	fmt.Println("valPtr=",valPtr,"&val=",&val)

	// Indirectly  Upadting  Value ,via Pointer

	*valPtr=999
	fmt.Println("val=", val,"valPtr=",*&valPtr)
	fmt.Println("valPtr=",valPtr,"&val=",&val)

	fmt.Println("val +10=",val+10)
	fmt.Println("*valPtr +10=",&valPtr +10)

	var val2 int =55
	var valPtr * int = &val2
	

	fmt.Println("val +val2 =",val +val2)
	fmt.Println("*valPtr + valPtr2 =",*valPtr + * valPtr)
	










}