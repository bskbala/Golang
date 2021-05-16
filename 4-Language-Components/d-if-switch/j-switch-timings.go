package main 

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("time.Now().Weekday()=",time.Now().Weekday())

	switch time.Now().Weekday() {  // Switch Initlzeer
	case time.Monday, time.Tuesday, time.Wednsday, time.Thrusday, time.Friday:
		fmt.Println("Office Timings:9am to 6 pm")
	case time.saturday:
		fmt.Println("office timings:9am to 1 pm ")
	case time.sunday:
		fmt.Println("====Hoilday===")
	default:
		fmt.Println("make this days are not happy")
	}
}