package main

import (
	"fmt"
)

func main (){
	x :=5
	fmt.Println("Before,x=",x, "&x=",x)

	zero(&x)
	fmt.Println("After,x=",x, "&x",&x)

}

func zero(z *int){
	fmt.Println("z=",z,"*z=",&z)
	*z =0
}