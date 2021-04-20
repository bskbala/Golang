package main

import "fmt"

func main(){

	fmt.Println("Please enter day of the week")

	var weekOfday string 
	fmt.Scanf("%s",&weekOfday)

	// Converting to lower cases 
	weekOfday = strings.ToLower(weekOfday)

	switch weekOfday{
	case "monday","tuesday","wednsday","thrusday","friday":
		fmt.Println("Office Timnings : 9 am to 6 pm")
	case "saturday":
		fmt.Println("Office Timnings : 9 am to 1 pm")
	case "sunday":
		fmt.Println("==== Today is Hoilday===")
	}
}