package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "i love food "
	half := len(s) / 2
	fmt.Println("s =", s)
	fmt.Println("s[:half] =", s[:half])

	result := strings.ToUpper(s[:half]) +  strings.ToLower(s[:half])
	fmt.Println(result)
	

}
