package main

import (
	"fmt"
	"time"
)

func main() {
	count :=0
	for range time.Tick(time.Second * 2){
		fmt.Println("Hello santra - %d\n",count)
		if count >5 {
			break
		}
		count +=1
	}
}