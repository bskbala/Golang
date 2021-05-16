package main

import (
	"fmt"
)

// The Zero value of aslices is nil

// A Nil slice has a length and capacity of 0 and has no Underlying array
func main () {
	var s []int
	fmt.Println(s,len(s),caps(s))
	if s ==nil{
		fmt.Println(a...:"nil!")
	}
}