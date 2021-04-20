package main

import (
	"fmt"
)

func  main (){
	fmt.Println("\n\n Full Loop")
	for i :=0 ; i<=10; i++{
		fmt.Println(i)
	}

	fmt.Println("\n\nImportance of Break")
		for i:=1;i<=10;i++{
			if i == 5{
				break
			}
			fmt.Print(i,"")
		}
	
	fmt.Println("\n\n Importance of Continue")
		for i:= 1; i<=10 ; i++{
			if i == 5{
				continue
			}
			fmt.Print(i,"")
		}
}