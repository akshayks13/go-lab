package main

import (
	"fmt"
)

func main(){
	var age int
	fmt.Println("Enter your age:")
	fmt.Scan(&age)

	// else must be on the same line as } 
	if age < 0 {
		fmt.Println("Invalid age entered.")
	} else if age<18{
		fmt.Println("You are a minor.")
	} else if age>=18 && age<65{
		fmt.Println("You are an adult.")
	}else if age>=65{
		fmt.Println("You are a senior citizen.")
	} else {
		fmt.Println("You entered a valid age:", age)
	}
}