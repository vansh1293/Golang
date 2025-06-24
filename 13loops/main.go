package main

import "fmt"

func main() {
	days := []string{"sunday", "monday", "tuesday", "wednesday"}
	fmt.Println(days)
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	for i := range days {
		fmt.Println(days[i])
	}
	for index, day := range days {
		fmt.Println(index, day)
	}
	value := 1
	for value <= 10 {
		if value == 2 {
			value++
			continue
		}
		if value == 3 {
			goto lco
		}
		if value == 5 {
			break
		}
		fmt.Println("Value is: ", value)
		value++
	}
lco:
	fmt.Println("Hello")
}
