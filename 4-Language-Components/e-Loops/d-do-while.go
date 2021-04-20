package main

import (
	"fmt"
)


// General Go don't support dO While
func main(){

	for ok := true ; ok; ok =false{
		fmt.Println("IN LOOP")
	}

	for  ok := true; ok;{
		fmt.Println("IN lOOP2")
		ok=false
	}
}
