package main

import (
	"fmt"
)

func main(){

	fmt.Println(new([]int))
	fmt.Println(make([]int,5))

	fmt.Println(new([]float64))
	fmt.Println(make([]float64,5))

	fmt.Println(new([]string))
	fmt.Println(make([]string,5))

	fmt.Println(new([]rune))
	fmt.Println(make([]rune,5))

	fmt.Println(new([]bool))
	fmt.Println(make([]bool,5))

	ci :=make(chan int)
	cs := make(chan string)
	cf :=make(chan interface{})

	fmt.Println("ci=",ci)
	fmt.Println("cs=",cs)
	fmt.Println("cf=",cf)

	
}