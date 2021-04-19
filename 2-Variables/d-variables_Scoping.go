package main

import "fmt"

var x string = "Globle Scope"

func main() {

	var x string = "block Scope "

		// Local will be preferred

		fmt.Println("In main () :", x, &x)

		myfunc()
		anotherfunc(x string)

}

func myfunc() {
	fmt.Println("In main() :" ,x, &x)
}

func anotherfunc(x string) {
	fmt.Println("In main() :" ,x, &x)
}
