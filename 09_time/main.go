package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("waiting for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Println("wait over!")

	var dateFormat string = "01-02-2006 Monday"

	presentTime := time.Now()
	fmt.Println("current time:", presentTime)
	fmt.Println("current time:", presentTime.Format(dateFormat))

	createdDate := time.Date(2020, time.July, 10, 23, 45, 0, 0, time.UTC)
	fmt.Println("createdDate:", createdDate.Format(dateFormat))
}
