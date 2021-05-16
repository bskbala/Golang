package main

import (
	"fmt"
)

func main(){
	var alphabets = map[string]string{
		"a":"apple",
		"b":"banana",
		"c":"cat",


	}
	fmt.Printf("alphabets=%v - %[1]T\n",alphabets)
	fmt.Println("refelct.TypeOf(alphabets),refelct.TypeOf(alphabets).Kind()")

	someThing :=map[int]bool{
	123:true,
	-234:false,
	34:true,
	5:true,
	


}
fmt.Println("someThing =",someThing)

}
