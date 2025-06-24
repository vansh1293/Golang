package main

import "fmt"

func main() {
	var fruits [4]string
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[3] = "pineapple"
	fmt.Println("Array is: ", fruits)
	fmt.Println("Array is: ", len(fruits))
	var veggies = [5]string{"potato", "tomato", "onion"}
	fmt.Println("Array is: ", veggies)
	fmt.Println("Array is: ", len(veggies))
}
