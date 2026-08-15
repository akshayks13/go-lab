package main

import (
	"fmt"
)

func main(){
	fmt.Println("hello world")

	var a = "Akshay"
	const b = 10 // We cant change the value of b
	fmt.Println("hello my name is ",a)
	fmt.Println("Number is ",b)
	fmt.Println(&a) // This will print the address of the variable a

	fmt.Printf("Hello %v, number is %d\n",a,b)  // %v is for any type, %d is for integer (can use %s for string)

	// If we want to decalre first and assigne later , we have to use the type we are going to use as well

	var name string
	var age int 
	name = "hello"
	age  = 20
	fmt.Println("My name is : " ,name)
	fmt.Println("My age is : " ,age)

	simply := "hi"
	fmt.Println("Simply is : ", simply) // This is a shorthand for variable declaration and assignment, it infers the type automatically
	// We can also use the short variable declaration inside a function
}
