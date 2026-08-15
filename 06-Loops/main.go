package main

import (
	"fmt"
)

// Go only has one loop construct: the `for` loop.
// It can be used in several ways, including as a traditional for loop, a while loop, or an infinite loop.

func main() {

	// Infinite loop example
	fmt.Println("Starting infinite loop...")
	for {
		fmt.Println("This will run forever unless you stop it manually.")
		break
	}
	fmt.Println("Exited infinite loop.")

	var arr = [5]int{1,2,3,4,5}
	for ind,val := range(arr) {
		fmt.Printf("Index: %d, Value: %d\n", ind, val)
	}

	// Traditional for loop example
	for i:=0; i<5; i++{
		fmt.Println("Value at index", i, "is", arr[i])
	}

	// While loop example
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}