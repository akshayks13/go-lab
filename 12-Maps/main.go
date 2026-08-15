package main

import "fmt"

func main(){
	// var m map[string]int // // declares a nil map ,But this is unusable until initialized.
	// m["key"] = 1 // This will cause a runtime panic: assignment to entry in nil map
	// m = make(map[string]int) // Initialize the map
	// m["key"] = 1 // Now we can assign a value to the map
	// fmt.Println("Map:", m) // Print the map

	var sampleMap = make(map[string]int) // Initialize the map
	sampleMap["one"] = 1
	sampleMap["two"] = 2
	sampleMap["three"] = 3

	delete(sampleMap, "two") // Delete the key "two" from the map
	fmt.Println("Sample Map:", sampleMap) // Print the sample map

	var arrOfMaps = make([]map[string]int ,3) // Initialize a slice of maps , the size of slice can be dynammically increased
	for i := 0; i < len(arrOfMaps); i++ {
		arrOfMaps[i] = make(map[string]int)
	}

	arrOfMaps[0]["one"] = 1
	arrOfMaps[0]["two"] = 2
	arrOfMaps[1]["two"] = 2
	arrOfMaps[2]["three"] = 3
	arrOfMaps[2]["four"] = 4
	fmt.Println("Array of Maps:", arrOfMaps) // Print the array of maps
}