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
	fmt.Println("User status:", Vansh.GetStatus())
	Vansh.NewEmail()                           // This will not change the original email bcoz copy of the object is passed
	fmt.Println("Email of user:", Vansh.Email) // Still the old email
}

// Capital first letter means public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() string {
	if u.Status {
		return "User is active"
	}
	return "User is inactive"
}
func (u User) NewEmail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of user is:", u.Email)
}
