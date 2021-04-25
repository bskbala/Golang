package main


import (
	"fmt"
)

func printASlice(s string ,x [] int){
fmt.Println("%s len =%d cap =%d %v\n",
	s,len(x),cap(x),x)
}


func main() {

	a :=make([])

}