package main

import (
	"fmt"
)

func main(){
	var name string
	var age int

	fmt.Print("Enter your name : ")
	fmt.Scan(&name) // We need to pass the address of the variable to store the input value
	fmt.Print("Enter your age : ")
	fmt.Scan(&age) 

	fmt.Println("Hello", name, "you are", age, "years old")
}