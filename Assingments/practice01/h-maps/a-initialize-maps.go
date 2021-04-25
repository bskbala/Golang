package Main

import (
	"fmt"
)

func main () {
	// Creating and Initializing empty map
	// Using avr Keyword

	var map_1 map[int]int

	// checking if he map is nil or not
	if map_1 == nil {
		fmt.Println("True")
	}else{
		fmt.Println("false")
	
	}
	// Creating anfd Initializing a map
	// Using shortand declaration and
	// using  map literals

	map_2 :=map[int]string{
		90 :"Dog",
		91 :"Cat",
		92 : "Cow",
		93: "Bird",
		94:"Rabbit",
	}
	fmt.Println("map-2:",map_2)
}