package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	defer fmt.Println("Deferred: This will run last")
	defer fmt.Println("Deferred: This will run second")
	defer fmt.Println("Deferred: This will run first")
	// The deferred functions are executed when the surrounding function returns
	// The deferred functions are executed in reverse order
	// The deferred functions will run in LIFO order
	// so the last deferred function will run first.
	// The output will be:
	// Deferred: This will run first
	// Deferred: This will run second
	// Deferred: This will run last
	fmt.Println("This will run first")
	myDefer()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("Deferred:", i)
	}
}
