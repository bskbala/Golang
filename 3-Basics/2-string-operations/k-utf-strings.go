package main

import (
	"fmt"
	"unicode/utf8"
	"golang.org/x/text/cases"
)

func main(){

	s := "Hello ,###"
	fmt.Println("string length:",len(s))
	fmt.Println("rune count :", utf8.RuneCountInStrings(s))

	s1 := "Hello ,world"
	fmt.Println("string length:",len(s1))
	fmt.Println("rune count :", utf8.RuneCountInStrings(s1))

	for i := 0; i < len(s) ; i++{
		fmt.Println(s[i:])
	}


}