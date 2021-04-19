package main

import (
	"fmt"
)

func main() {
	fmt.Println("\t rune : %d \n", 'a')
	fmt.Println("\t rune : %q \n", 'a')
	fmt.Println("\t UTF8 code :%v \n", 'a')
	fmt.Println()

	fmt.Println("\t rune : %c \n", 'z')
	fmt.Println("\t rune : %q \n", 'z')
	fmt.Println("\t rune : %a \n", 'z')
	fmt.Println()

	s := "hello world"
	encryptedString := ""
	for i := 0; i < len(s); i++ {
		// fmt.Println("\t rune : %c \t", s[i])
		// fmt.Println("\t UTF8 code : %v \t", s[i])

		// // fmt.Println("\t rune : %c \t" , s[i])

		// // Caesar Cipher :Charcater +3

		// fmt.Println("\t rune :%c \t", s[i]+3)
		// fmt.Println("\t UTF8 code : %v \n\n", s[i]+3)

		// encryptedString += s[i] +3
		encryptedString += fmt.Sprintf("%c", s[i]+3)
		fmt.Println()
	}
	fmt.Println("encryptedString =", encryptedString)

}
