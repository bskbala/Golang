package main

import "fmt"

func main() {

	fmt.Println(`len (Hello world)= `, len("hello world"))

	name := "Hello World"
	fmt.Println("name  =", name)
	fmt.Println("len(name)  =", len(name))

	fmt.Println()
	fmt.Println("name[0]   =", name[0])
	fmt.Println("name[0]   = %c\n", name[0])
	fmt.Println("name[4]   = %c\n", name[4])
	fmt.Println("name[5]   = %q\n", name[5])
	fmt.Println("name[10]  = %q\n", name[10])

	// Slicing -last position  char is not Included

	fmt.Println("name[2:7]   =", name[2:7])
	fmt.Println("name[1:10]   =", name[1:10])
	fmt.Println("name[1:11]   =", name[1:11])

	fmt.Println()
	fmt.Println("name[1:]  =", name[1:])
	fmt.Println("name[:]=", name[:])

	fmt.Println()
	fmt.Println("name[len(name) -3:]=", name[len("name")-3:])
	fmt.Println("name[len(name) -3:]=", name[len("name")-3:])

}
