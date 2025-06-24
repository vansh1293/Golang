package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Go")
	presentTime := time.Now()
	fmt.Println("Present Time is: ", presentTime)
	fmt.Println("Present Time is: ", presentTime.Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2005, time.March, 11, 7, 30, 0, 0, time.UTC)
	fmt.Println("Created Date is: ", createdDate)
	fmt.Println("Created Date is: ", createdDate.Format("01-02-2006 Monday"))
}
