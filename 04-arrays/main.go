package main

import "fmt"

func main(){
	var arr = [5]int{1,2,3} // Declare an array of 5 integers or even var arr [5]int{1,2,3}
	arr[3] = 4 
	arr[4] = 5 

	fmt.Println("Array:", arr) // Print the array

	var arr2 [5]string // if we dont initialize the array, it will be filled with empty strings
	arr2[0] = "Hello"
	arr2[1] = "World"
	arr2[2] = "Go"
	// arr2[3] = "is"
	// arr2[4] = "awesome"
	fmt.Println("String Array:", arr2)
	fmt.Println("Length of arr2:", len(arr2)) // Print the length of the array
	fmt.Println("Capacity of arr2:", cap(arr2)) // Print the capacity of the array	
	// Note: In Go, the length and capacity of an array are the same, but for slices, they can be different.
}