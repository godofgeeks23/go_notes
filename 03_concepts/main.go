package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum", sum)
	fmt.Println("sum squared = ", pow(float64(sum), 2, 99999))

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite loop
	// for {
	// }

	// Like for, the if statement can start with a short statement to execute before the condition

}
