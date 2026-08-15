package main

import (
	"fmt"
	"time"
	"sync"
)

func printDetails( wg *sync.WaitGroup,name string,age int) {
	time.Sleep(2 * time.Second) // Simulating a delay
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	wg.Done() // Signal that this goroutine is done

	// defer wg.Done() // Ensure that Done is called when the function exits, even if it panics -- Recommended
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // We are going to run 2 goroutines

	go printDetails(&wg,"Akshay", 20)
	go printDetails(&wg,"John", 25)

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines finished executing.")
}