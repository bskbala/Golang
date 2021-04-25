package main

import (
	"fmt"
)

func main () {
	
	m_a_p :=map[int]string{
		90:"Dog",
		91:"Cat",
		92:"Cow",
		93:"Bird",
		94:"Rabbit",

	}
	// Orignal Map
	fmt.Println("\nOrginal map:",m_a_p)

	//  Adding  new key values pairs in the  map

	m_a_p[95] = "Parrot"
	m_a_p[96] = "Crow"
	fmt.Println("Map After adding new Key-Values pair :\n",m_a_p)

	// Updating the Key-values pairs in the map

	m_a_p[91] = "Pig"
	m_a_p[92] = "santra"
	fmt.Println("\nMap After updating  new Key-Values pair :\n",m_a_p)

}