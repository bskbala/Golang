package main

import  (
	"fmt"
)

func main (){
	var n,sum  int
	fmt.Println("Enter Postive Number")
	fmt.Scanf("%d",&n)
	for i :=1 ; i<=n ; i++{
	sum +=i

	}
	fmt.Println("sum",sum)

}
