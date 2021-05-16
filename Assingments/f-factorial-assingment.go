package main

import (
	"fmt"
)

// Variable Declaration
var factVal uint64 =1
var i int  =1
var n int
func factorial(n int) uint64{
	if (n<0){
		fmt.Print("factorial of negative Number Doesnt't Exist")

	}else{
		for i:=1; i<=n; i++{
			factVal *= uint64(i)
		}
	}
	return factVal
}
func main(){
	 fmt.Print("Enter a postive Interger Bettween 0-50:")
	 fmt.Scan(&n)
	 fmt.Print("factorial is:" ,factorial(n))
	
}