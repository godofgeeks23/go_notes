package main

import (
	"fmt"
)

func main() {
	fmt.Println("loops")

	labelCount := 1

label:
	days := []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
	fmt.Println(days)

	// in go we just have for loop. we also do have break and continue
	for d := 0; d < len(days); d++ {
		fmt.Println(d, days[d])
	}

	for d := range days {
		fmt.Println(d, days[d])
	}

	for index, day := range days {
		fmt.Println(index, day)
	}

	if labelCount < 2 {
		labelCount++
		goto label
	}
}
