package main

import (
	"fmt"
)

func main () {

	garr1 :=[]int{10,29,70,40,127}
	garr2 :=[]int{15,25,35,45,55}

	fmt.Println("Addition\tsubtraction\tMultiplication\tDivision\tModules")
	for i:=0; i<5; i++ {
		fmt.Println("\n",garr1[i]+garr2[i],"\t")
		fmt.Print(garr1[i]-garr2[i], "\t")
		fmt.Print(garr1[i]*garr2[i], "\t")
		fmt.Print(garr1[i]%garr2[i], "\t")
		fmt.Print(garr1[i]/garr2[i], "\t")
		

	}
	fmt.Println()

}