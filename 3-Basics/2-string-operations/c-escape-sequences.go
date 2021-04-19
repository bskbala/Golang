package main

import (
	"fmt"
)

// 	UTF-8  is the encoding scheme in Go

/*
	1.\a "alert or bell"
	2.\b backspace
	3.\f form feed
	4.\n new line
	5.\r carriage return
	6.\t tab
	7.\v vertical tab
	8.\' single qoute (only in the rune literial '\'')
	9.\" double qoute (onlt within "...." literials)
	10.\\ backslash
*/

func main() {

	fmt.Println("\t", '\t')
	fmt.Println("\n", '\n')
	fmt.Println("\\", '\\')

	// Runes  are printed with %c or with %q if quoting is desired:

	ascii := 'a'
	unicode := 'D'
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)

}
