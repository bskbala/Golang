package main

import (
	"fmt"
	"string"
)

func main (){
	println("Enter day of the week:")

	var weekOfday string
	fmt.Scanf("%s",&weekOfday)

	weekOfday =strings.ToLower(weekOfday)

	if weekOfday == "monday" ||  weekOfday == "Tuesday" || weekOfday == "Wednsday" ||  weekOfday == "Thursday" ||  weekOfday == "friday"{
		println("Office Timings :9 am to 6 pm ")
	}else if weekOfday == "saturday"{
		println("Office Timings :9 am to 1 pm ")
	}else if weekOfday == "sunday"{
		println("======Hoilday======= ")
	}
	}

