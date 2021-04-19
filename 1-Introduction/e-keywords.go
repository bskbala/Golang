package main

import "fmt"

func main() {

	fruit := "apple"
	fmt.Println("Fruit:", fruit)

	//  Note 1: keywords should not be used as identifiers
	// break := "one"  // syntax error : Unexpected :=at end of statement

	// Not Recommended - predeclareed varaiables Sho8uld not be used as keywords

	true := "one"
	fmt.Println("true:", true)

	break_one := "one"
	fmt.Println("break_one=", break_one)

	//  Note 2: Camelcase is recommended identifiers

	breakone :="one"
	fmt.Println("breakone=",breakone)

	costOfMangos := 34
	fmt,Println("costOfMangos =", costOfMangos)

	NoOfProcessesRunning :=3
	fmt.Println("NoOfProcessesRunning")

}
