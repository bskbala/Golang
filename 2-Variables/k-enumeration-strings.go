package main

import "fmt"

// data type aliases

type Direction int

func main() {
	const (
		North Direction = iota
		East
		south
		west
	)
	fmt.Println("south =", south)
	fmt.Println("NorthEast=", North, East)
}
