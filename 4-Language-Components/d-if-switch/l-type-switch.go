package main

import ( 
	"fmt"
)

func main (){
	getDataType(true)
	getDataType(1)
	getDataType("Hello")
	getDataType('h')
	getDataType(123.213)
}

func  getDataType(value interface{}){
	fmt.Printf("value =%v \t dataType -%T \n",value,value)

	switch value.(type){
	case bool:
		fmt.Println("I'm a bool")
	case int:
		fmt.Println("I'm an int")
	default:
		fmt.Println("Invalid data as Nibbas")
	}

}