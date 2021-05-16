package main

import (
	"fmt"
	"strings"
)

func main (){
	println("Enter day of the week:")

	var weekOfday string
	fmt.Scanf("%s",&weekOfday)

	weekOfday = strings.ToLower(weekOfday)
	if weekOfday == "monday"{
		println("Office Timings :9 am to 6 pm ")
	} else if weekOfday == "Tuesday"{
		println("Office Timings :9 am to 6 pm ")
	} else if weekOfday == "wednsday"{
		println("Office Timings :9 am to 6 pm ")
	}else if weekOfday == "Thursday"{
		println("Office Timings :9 am to 6 pm ")
	}else if weekOfday == "friday"{
		println("Office Timings :9 am to 6 pm ")
	}else if weekOfday == "saturday"{
		println("Office Timings :9 am to 1 pm ")
	}else if weekOfday == "sunday"{
		println("======Hoilday======= ")
	}


}