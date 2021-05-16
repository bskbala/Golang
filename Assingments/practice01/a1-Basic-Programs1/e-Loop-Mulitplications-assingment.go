package main

import ( 
	"fmt"
)

func main(){

	// for i :=1; i<=10 ;i++{
	// 	fmt.Println(i)
	// }

	// for j :=1; j<=10 ;j++{
	// 	fmt.Println(j)
	// }

	
	for i :=10; i>=1 ;i--{
		for j:=10; j>=1 ;j--{

			fmt.Println(i, "*", j, "=",  i * j)
			
	}
	
		 fmt.Println()
	}

}