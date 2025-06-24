package main

import (
	"fmt"
	"reflect"
)

func main() {
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("Type of languages is: ", reflect.TypeOf(languages))
	fmt.Println("Value of languages is: ", languages)
	fmt.Println("JS short for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("Value of languages is: ", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
