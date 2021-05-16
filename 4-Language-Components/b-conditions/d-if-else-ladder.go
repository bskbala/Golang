package main

import (
	"fmt"
)


/*
	+90 North Pole
	Nothern Hemisphere
	0 Equator
	south Hemisphere
	-90 South pole
*/
func main (){

	var lattitude float32
	fmt.Println("Please enter the lattiude degrees:")
	fmt.Scanf("%f",&lattitude)

	if lattitude == 90{
		fmt.Println("your are located at North Pole")
	}else if lattitude < 90 && lattitude >0 {
		fmt.Println("your are located in North Hemisphere")
	}else if lattitude == 0 {
		fmt.Println("youtr located  in Equator")
	}else if lattitude < 0 && lattitude >-90{
		fmt.Println("your are located at southern Hemisphere")
	}else if lattitude == -90{
		fmt.Println("your are located in the south Pole")
	}else {
		fmt.Println("you have entered Danger lattitude Zone")
	}


}