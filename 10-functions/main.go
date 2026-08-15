package main

import (
	"fmt"
)

func main(){
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name) 
	if name == "" {
		fmt.Println("Name cannot be empty. Please enter a valid name.")
		return
	}

	welcome(name)

	var a int
	var b int
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	fmt.Println("Sum:", add(a, b))
	fmt.Println("Difference:", subtract(a, b))

	fmt.Print("Enter numbers of elements in the array: ")
	var n int
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Invalid number of elements. Please enter a positive integer.")
		return
	}
	var arr []int
	for i:=0; i<n; i++ {
		var num int
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scan(&num)
		arr = append(arr, num)
	}

	var sum int = sumofarray(arr)
	fmt.Println("Sum of array elements:",arr," is ",sum)

}

// Note: Go functions can return multiple values.

func welcome(name string){
	fmt.Println("Welcome", name)
}

func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}

func sumofarray(arr []int)int{
	var sum int = 0 
	for _,num := range arr{
		sum+=num
	}
	return sum
}