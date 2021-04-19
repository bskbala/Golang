package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main(){

	x :=123
	y := fmt.Sprintf("num-%d",x)
	fmt.Println("x=",x,reflect.TypeOf(x))
	fmt.Println("y=",y,reflect.TypeOf(y))
	fmt.Println()

	z := strconv.Itoa(x)
	fmt.Println("z=",z,reflect.TypeOf(z))
	fmt.Println()

	// To formate number in Different base

	fmt.Println(strconv.FormatInt(int64(x),2))
	fmt.Println(strconv.FormatInt(int64(x),10))
	fmt.Println(strconv.FormatInt(int64(x),16))
	s :=fmt.Sprintf("x=%b",x)
	fmt.Println(s)
	fmt.Println()


	// To format string  represting an integer

	x,err := strconv.Atoi("123")
	fmt.Println(x,err,reflect.TypeOf(x))
	fmt.Println()

	// p.err1 :=strconv.ParseInt("123312.123, 10, 64")
	// fmt.Println(p, err1)

	

}