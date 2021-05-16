package main

import (
	"fmt"
	"time"
)

func main (){

	today :=time.Now()
	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th .Clean  your House")
		fallthrough
	case 10:
		fmt.Println("Today is 10th .Clean  your clothes")
		fallthrough
	case 12:
		fmt.Println("Today is 12th .Go for Shopping ")
		fallthrough	
	case 14:
		fmt.Println("Today is 14th.Make some cake for Kids")
		
	default:
		fmt.Println("No information is need for the day ")
	}
}


