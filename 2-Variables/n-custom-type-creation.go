package main

import (
	"fmt"
	"reflect"
)

type fruit int

func main() {

	var apple int = 10
	fmt.Println(apple, reflect.TypeOf(apple))

	var banana fruit = 10
	fmt.Println(banana, reflect.TypeOf(banana), reflect.TypeOf(banana).kind())

}
