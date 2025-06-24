package main

import "fmt"

func main() {
	var ptr *int // if we don't initialize the pointer it will be nil
	fmt.Println("Value of pointer is: ", ptr)
	num := 23
	var ptr1 = &num // pointer is address to direct memory location
	fmt.Println("Value of pointer is: ", ptr1)
	fmt.Println("Value of pointer is: ", *ptr1)
	*ptr1 = *ptr1 * 2
	fmt.Println("Value of pointer is: ", num)
}
