package main

import (
	"fmt"
)

// Declaring Package Level  

const LUCKYNUMBER int32 =67
func main(){
	fmt.Println("Guess a number btwn 1 and 100")

	var guessedNumber int32
	fmt.Scanf("%d",&guessedNumber)

	// Logic 1
	if guessedNumber == LUCKYNUMBER {
		println("you guessed correctly")
	}

	// Logic 2

	if guessedNumber == LUCKYNUMBER {
		println("You gussed Correctly")
	}else{
		println("please Try Again")
	}

	// Logic 3

	if guessedNumber == LUCKYNUMBER {
		println("You gussed Correctly")
	}else if guessedNumber > LUCKYNUMBER {
		println("please Reduce your Guessing Number ")
	}else {
		println("please increase your guessing")
	}



}

