package main

import ( "fmt"
)

func main (){

	fmt.Println("Enter first Number")
	var first int
	fmt.Scanln(&first)
	fmt.Println("Enter second  Number:")
	var second int
	fmt.Scanf("%d" ,&second)
	fmt.Println("Enter third  Number:")
	var third int
	fmt.Scanf("%d" ,&third)

	if (first >=second && first >=third){
		fmt.Println("first is the largest Numeber")
	}else if (second >=first && second >=third){
		fmt.Println("second is the largest Numeber")
	}else if (third >=first && third >=second){
		fmt.Println("third is the largest Numeber")
	}

}