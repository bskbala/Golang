package main

import (
	"fmt"
)

func main (){
	var n,sum int
	fmt.Print("Enter a Postive integer :")
	fmt.Scan(&n)
	for i:=1; i<=n; i++{
		sum +=i
	}
	fmt.Print("Sum:",sum)
}