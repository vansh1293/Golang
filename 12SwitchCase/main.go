package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn((6)) + 1
	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 and you can move 2 spot")
	case 3:
		fmt.Println("Dice value is 3 and you can move 3 spot")
	case 4:
		fmt.Println("Dice value is 4 and you can move 4 spot")
		// fallthrough - adding this will do that if we hit this case and this case and all the cases below it will be executed
	case 5:
		fmt.Println("Dice value is 5 and you can move 5 spot")
	case 6:
		fmt.Println("Dice value is 6 and you can move 6 spot and roll dice again")
	default:
		fmt.Println("You got a snake eye")
	}
}
