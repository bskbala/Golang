package main

import (
	"fmt"
)

func main(){

	switch days:=2;days{ //Note :;semicolon is statement Seprator
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
		fallthrough
	case 4:
		fmt.Println("Four")
		// fallthrough
	default:
		fmt.Println("NIBBA")

	}
	// FallThrough Statmment will Be executed

}