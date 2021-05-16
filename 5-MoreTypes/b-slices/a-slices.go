package main

import (
	"fmt"
	"reflect"
)


func main() {

	a1 := [...]string{"a","b","c","d","e","f","g"}

	// Last Index is not Included,In slices

	s1 :=a1[2:4]
	s2 :=a1[3:6]
	s3 :=a1[1:5]
	s4 :=a1[:]
	s5 :=a1[:7]

	fmt.Println("\n a1=", a1, reflect.TypeOf(a1).Kind())
	fmt.Println("len(a1))=",len(a1))
	fmt.Println("cap(a1))=",cap(a1))

	fmt.Println("\n s1=", a1, reflect.TypeOf(s1).Kind())
	fmt.Println("len(s1))=",len(s1))
	fmt.Println("cap(s1))=",cap(s1))

	fmt.Println("\n s2=", a1, reflect.TypeOf(s2).Kind())
	fmt.Println("len(s2))=",len(s2))
	fmt.Println("cap(s2))=",cap(s2))

	fmt.Println("\n s3=", a1, reflect.TypeOf(s3).Kind())
	fmt.Println("len(s3))=",len(s3))
	fmt.Println("cap(s3))=",cap(s3))

	fmt.Println("\n s4=", a1, reflect.TypeOf(s4).Kind())
	fmt.Println("len(s4))=",len(s4))
	fmt.Println("cap(s4))=",cap(s4))

	fmt.Println("\n s5=", a1, reflect.TypeOf(s5).Kind())
	fmt.Println("len(s5))=",len(s5))
	fmt.Println("cap(s5))=",cap(s5))



	
	
}