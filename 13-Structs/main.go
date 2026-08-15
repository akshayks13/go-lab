package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main(){
	var p = Person{
		Name: "Akshay",
		Age:  20,
	}

	fmt.Println("Name:", p.Name)

	p.Age = 19
	fmt.Println("Age:", p.Age)
}