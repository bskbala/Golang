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

	// Reterving Values with the help of keys
	value_1 := m_a_p[90]
	value_2 := m_a_p[93]
	fmt.Println("Value of Key [90]:",value_1)
	fmt.Println("Value of Key [93]:",value_2)


}