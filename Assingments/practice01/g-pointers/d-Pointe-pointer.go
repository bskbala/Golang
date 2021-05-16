package main


import (
	"fmt"
)
// Main Function
func main() {

	// taking a varaible
	// of integer Type
	var v int=100

	// taking a pointe
	// of integer type.
	var  pt1 * int =&v

	// taking pointer to
	// pointer to pt1
	// storing the address
	// of pt1 into pt2
	var pt2 ** int =& pt1

	fmt.Println("The value of varaible v is =",v)

	// changing the v by assiging
	// the new value to the pointer pt1

	*pt1 =200

	fmt.Println("value stored in v aftre changing pt1 =",v)

	// changing the value of v by assiging
	// the new value to theh Pointer pt2

	**pt2 = 300

	fmt.Println("value stored in v aferr changing pt2 =",v)





}