package main

import "fmt"

func main() {
	emptyString1 := ""
	var emptyString2 string // default   Intaila value for the strinf is ""
	var emptyString3 = ""
	var emptyString4 string = ""
	const emptyString5 = ""

	fmt.Println("emptyString1 == emptyString 2 :", emptyString1 == emptyString2) //true
	fmt.Println("emptyString1 == emptyString 3 :", emptyString1 == emptyString3) //true
	fmt.Println("emptyString1 == emptyString 4 :", emptyString1 == emptyString4) //true
	fmt.Println("emptyString1 == emptyString 5 :", emptyString1 == emptyString5) //true

	// TO get the address location where an object is stored
	fmt.Println("emptyString1 == emptyString 1 :", emptyString1 == emptyString1)  //true
	fmt.Println("emptyString2 == emptyString 2 :", emptyString2 == &emptyString2) //true
	fmt.Println("emptyString3 == emptyString 3 :", emptyString3 == &emptyString3) //true
	fmt.Println("emptyString4 == emptyString 4 :", emptyString4 == &emptyString4) //true

}
