package main

import (
	"fmt"
)

func main (){

	fmt.Println("In main() loop -start")

}

func myfunc(){
	fmt.Println("in myfunc - start")

}

func init(){
	fmt.Println("In init() function-start")
}