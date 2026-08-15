package main

import (
    "fmt"
)

func sayHi(chRecv chan string, chSend chan string) {
    msg := <-chRecv                // receive message from main
    fmt.Println("Goroutine received:", msg)
    chSend <- "Hi from goroutine!" // send response back
}

func main() {
    chToGoroutine := make(chan string)
    chFromGoroutine := make(chan string)

    go sayHi(chToGoroutine, chFromGoroutine)

    chToGoroutine <- "Hello from main!"         // send message to goroutine
    msg := <-chFromGoroutine                    // receive response from goroutine
    fmt.Println("Main received:", msg)
}