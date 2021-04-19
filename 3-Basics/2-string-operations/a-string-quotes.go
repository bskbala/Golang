package main

import "fmt"

/* 
 Go natively handels Unicode,so 
 it can process twxt in all the world"languages

*/

func main(){

	fmt.Println("H")
	fmt.Println('H')
	fmt.Println("")

	fmt.Println("Hello World")

	// double Qoute Strings -Single Line Strings

	fmt.Println("Hello go World")
	fmt.Println("Hello\tgo\nworld")
	fmt.Println("Hello\bgo\nworld")
	fmt.Println("123\bDog")
	fmt.Println("123456\rDog")

	fmt.Println("Quotes:'")
	fmt.Println("Qoutes: \"")
	fmt.Println("Backlash:\\")

	// Multil line Strings

	// fmt.Println("Hello 
	// 				World")


	// Back tick(raw) strings -Both Single & Multiple

	fmt.Println(`H`)
	fmt.Println(`hello go world`)
	fmt.Println(`hello\tog\nworld`)
	fmt.Println(`quotes:"`)

	fmt.Println(`Hello 
		World`)

}