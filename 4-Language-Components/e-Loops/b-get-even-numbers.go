package main

import (
	"fmt"
)

func main(){
	fmt.Println("Please Print Even Numbers 0 & 255:")

	for i:=0 ;i<=255;i++{
		if i%2 ==0{

			fmt.Printf("%d \t",i)
		}

	}
}