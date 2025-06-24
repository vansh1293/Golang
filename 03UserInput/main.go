package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	welcome := "Hello"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || error ok
	input, _ := reader.ReadString('\n') // we can do _,error or input,error also
	fmt.Println("Thanks for rating, ", input)
	fmt.Println("Type of this rating is: ", reflect.TypeOf(input))
}
