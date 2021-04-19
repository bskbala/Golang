package main

import (
	"fmt"
)

func main (){
	mystr :="Hello world"
	fmt.Println("mystr =", mystr)

	fmt.Println("mystr[6:11] =", mystr[6:11])

	myrune := 'H'
	fmt.Println("before myrune =", myrune)

	myrune = 'W'
	fmt.Println("After myrune =", myrune)




}
