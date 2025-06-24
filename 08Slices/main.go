package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	var fruits = []string{"apple", "banana", "pineapple", "mango"}
	fmt.Println("Type of fruits is: ", reflect.TypeOf(fruits))
	fruits = append(fruits, "kiwi", "guava")
	fmt.Println("Fruits are: ", fruits)
	fruits = append(fruits[1:]) //tells start from 1 to last and if we define 1:3 then it will print from 1 to 2 index
	// [:3] tells start from 0 to 2 index
	fmt.Println("Fruits are: ", fruits)
	// there are two keywords to allocate memory make() and new().
	highScore := make([]int, 4) // this will allocate memory for 4 elements
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	fmt.Println("High Score are: ", highScore)
	fmt.Println("High Score are: ", len(highScore))
	fmt.Println("High Score are: ", cap(highScore)) // capacity of the slice
	highScore = append(highScore, 555, 666, 777)    // here memory allocation happens again
	fmt.Println("High Score are: ", highScore)
	fmt.Println("High Score are: ", len(highScore))
	fmt.Println("High Score are: ", cap(highScore))
	sort.Ints(highScore)
	fmt.Println("High Score are: ", highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	var course = []string{"reactjs", "javascript", "python", "java", "c++", "c#"}
	fmt.Println("Course is: ", course)
	var index int = 2
	course = append(course[:index], course[index+1:]...) // here we are removing the element from the slice
	fmt.Println("Course is: ", course)
}
