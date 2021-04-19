package main

import  "fmt"

func main() {

	var X ,Y int
	fmt.Println("\n Intially ,X: %3d , Y:%d",X,Y)

	fmt.Sscan("100\n200",&X,&Y)
	fmt.Println("\nSscan,X: %3d , Y:%d",X,Y)

	fmt.Sscan("100.23\n200.4",&X,&Y)
	fmt.Println("\nSscan, X: %3d , Y:%d",X,Y)
}