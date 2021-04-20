package main

import (
	"fmt"
	"time"
)

func main (){
	t := time.Now()
	fmt.Println(t.Hour())
	// TagLess Switch
	switch{
		case t.Hour() <=12:
			fmt.Println("Good Morning")
		case t.Hour() < 17:
			fmt.Println("Good AfterNoon")
		default:
			fmt.Println("Good Evening")	
	}

}
