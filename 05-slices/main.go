package main

import "fmt"

func main() {
	var slice []int // Declare a slice of integers
	slice = append(slice, 1) // Append 1 to the slice
	slice = append(slice, 2) // Append 2 to the slice

	fmt.Println("Slice:", slice) // Print the slice
	slice = append(slice, 3) // Append 3 to the slice
	fmt.Println("Slice after appending 3:", slice) // Print the slice after appending 3
}