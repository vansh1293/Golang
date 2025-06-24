package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
func farewell(name string) {
	fmt.Printf("Goodbye, %s!\n", name)
}
func add(a int, b int) int {
	return a + b
}
func proadder(values ...int) (int, string) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum, "hehehe"
}
func main() {
	fmt.Println("Welcome to the Go Functions Example!")
	greet("Alice")
	farewell("Alice")
	result := add(5, 10)
	fmt.Printf("The sum of 5 and 10 is: %d\n", result)
	proadderResult, str := proadder(1, 2, 3, 4, 5)
	fmt.Printf("The sum of 1, 2, 3, 4, and 5 is: %d\n", proadderResult)
	fmt.Printf("The string returned by proadder is: %s\n", str)
}
