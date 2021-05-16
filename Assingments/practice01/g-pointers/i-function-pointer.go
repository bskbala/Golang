package main 

import
	"fmt"


// taking a function  with integer
// type pointer as ann parameter
func ptf(a *int) {
	// dereferencing
	*a =748
}
func main() {
	
	// taking a normal variable

	var x= 100
	 fmt.Printf("The value of x before Function  call is :%d\n",x)

	//  Taking a pointer variable
	//  and assiging the address
	// of x to it 

	var pa * int =&x

	// calling the function by 
	// passing Pointer to function
	ptf(pa)
	fmt.Printf("the value of x after funcation call: %d\n",x)


}