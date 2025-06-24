package main

import "fmt"

const LoginToken string = "token" // we make this variable public by capitalizing the first letter and now we
// can use it in other files in the same folder or in this file

func main() {
	fmt.Println("Variables")
	var username string = "Vansh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)
	var smallVal uint8 = 255 // 256 is out of bound in uint8
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)
	var smallFloat float32 = 255.12345896 //in float32 we only 5 values after decimal - will print 255.12346
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	// Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	var anotherVariable2 string
	fmt.Println(anotherVariable2)
	fmt.Printf("Variable is of type: %T\n", anotherVariable2)

	// Implicit type
	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T\n", website)
	// lexar automatically predcits the type based on the value
	// variable website has string type not so we cannot change it further

	// no var style
	numberOfUsers := 3000000 // we cannot use this syntax for global variable
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T\n", numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)
}
