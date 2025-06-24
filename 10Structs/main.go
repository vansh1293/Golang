package main

import "fmt"

func main() {
	// no inheritance
	// no method overloading
	// no super or parent or child
	// no class
	Vansh := User{"Vansh", "vansh@go.dev", true, 20}
	fmt.Println(Vansh)
	fmt.Printf("Vansh details are: %+v\n", Vansh)
	fmt.Printf("Name is: %v\n", Vansh.Name)
}

// Capital first letter means public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
