package main

import (
	"fmt"
	"time"
)

func main(){
	today :=time.Now()
	var t int = today.Day()
	switch t {
	case 5:
		fmt.Println("Today is 5th .Clean  your House")
	case 10:
		fmt.Println("Today is 10th .Clean  your clothes")
	case 12:
		fmt.Println("Today is 12th .Go for Shopping ")	
	case 14:
		fmt.Println("Today is 14th.Make some cake for Kids")
	}
}