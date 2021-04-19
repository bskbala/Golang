// khoor#zruog

package main

import (
	"fmt"
)

func main() {

	s := "khoor#zruog"
	dcryptedString := ""
	for i := 0; i < len(s); i++ {
		// fmt.Println("\t rune : %c \t", s[i])
		// fmt.Println("\t UTF8 code : %v \t", s[i])

		// // fmt.Println("\t rune : %c \t" , s[i])

		// // Caesar Cipher :Charcater -3

		// fmt.Println("\t rune :%c \t", s[i]-3)
		// fmt.Println("\t UTF8 code : %v \n\n", s[i]-3)

		// DcryptedString += s[i] +3
		dcryptedString += fmt.Sprintf("%c", s[i]-3)
		fmt.Println()
	}
	fmt.Println("dcryptedString =", dcryptedString)
}
