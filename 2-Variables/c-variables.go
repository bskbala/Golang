package main

import "fmt"

var num1, num2 int = 1, 3

// Expotable varaibles Should be Seprately  Declared
var Num2 int = 2

func main() {

	// Local Scope  Variable
	num1 := "one"

	// Multiple type values interface

	var lang, version, islatest = "go", 1.15, true
	fmt.Println("num1 :", num1)
	fmt.Println("num2 :", num2)
	fmt.Println("lang :", lang)
	fmt.Println("version :", version)
	fmt.Println("islatest :", islatest)

}

/*
	Note:
	1.Unused Gobale Variables will be Ignored.
	No Compilation Errors
	2.Variables are of two types
	 1.1.Publical Exportable Variabels
	   - all variables starting with capital letter
	   - all unicode(non -English) variables

	 1.2.Private Variables
	   -Variables starting with small case Letter
*/
