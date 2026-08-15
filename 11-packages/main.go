package main

import "fmt"

func main(){
	var n int
	fmt.Print("Enter a number to calculate its factorial: ")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("Factorial is not defined for negative numbers.")
		return
	}
	result := factorial(n)
	fmt.Printf("Factorial of %d is %d\n", n, result)
}


// When we are using multiple modules, we need to import them.
// In this case, we are importing the "fact" package which contains the factorial function.

// If we want to import a package from a different directory, we need to specify the path.
// The function first letter must be capitalized to be exported.