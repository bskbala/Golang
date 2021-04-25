package main

import (
	"fmt"
)

func main (){
	months :=[...]string{
		1:"January"
		2:"February"
		3:"March"
		4:"April"
		5:"May"
		6:"June"
		7:"July"
		8:"August"
		9:"September"
		10: "Octomber"

	}

	Q2 := months[4:7]
	fmt.Println(a..."Q2 :" ,Q2)

	summer := months[6:9]
	fmt.Println(a..."summer:","summer")
}