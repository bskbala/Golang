package main 

import (
	"fmt"
)
// Range can be used to iterate(loop) over any iterable object
func main() {
	name := "GoLanguage"
	for i:=0 ; i<=len(name) ;i++{
		fmt.Printf("%q\n",name[i])
	}

// Method -2
for index := range name{
	fmt.Println("index=",index)

}
}