package  main


import (
	"fmt"
)

func main(){
	var stack []string 

	// Pushing
	stack = append(stack,"one")
	stack = append(stack,"Two")
	stack = append(stack,"Three")

	// Display

	fmt.Println("stack =" ,stack)
}