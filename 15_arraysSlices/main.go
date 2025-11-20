package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("arrays")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "tomato"
	fruitList[3] = "peach"
	fmt.Println(fruitList, len(fruitList))

	var vegList = [3]string{"potato", "peas", "onion"}
	fmt.Println(vegList)

	// slices are under the hood arrays
	var fruitSlice []string // []
	fmt.Println(fruitSlice)

	var vegSlice = []string{"onion", "peas", "cabbage"}
	vegSlice = append(vegSlice, "broccoli", "spinach")
	fmt.Println(vegSlice)

	vegSlice = append(vegSlice[:3])
	fmt.Println(vegSlice)

	vegSlice = append(vegSlice[1:3]) // last range is not inclusive
	fmt.Println(vegSlice)

	highScores := make([]int, 4)
	highScores[0] = 678
	highScores[1] = 334
	highScores[2] = 665
	highScores[3] = 349
	// highScores[4] = 1983                           // out of bounds
	highScores = append(highScores, 555, 666, 777) // this will work as it reallocates the memory
	sort.Ints(highScores)
	fmt.Println(highScores)

	var indexToRemove int = 2
	highScores = append(highScores[:indexToRemove], highScores[indexToRemove+1:]...)
	fmt.Println(highScores)
}
