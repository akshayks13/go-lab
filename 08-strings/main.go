package main

import (
	"fmt"
	"strings"
)

func main(){
	var age int 
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name) 
	if (strings.ContainsAny(name ,"0123456789!#$%^&*()_+={}[]|\\:;\"'<>,.?/~`-")) || (len(name) <= 2) {
		fmt.Println("Invalid name entered. Please enter a valid name without special characters or numbers.")
		return
	}

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	if age < 0 {
		fmt.Println("Invalid age entered. Age cannot be negative.")
		return
	} else if age > 120 {
		fmt.Println("Invalid age entered. Age cannot be greater than 120.")
		return
	}
}