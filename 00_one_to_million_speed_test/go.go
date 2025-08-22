package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {

	writer := bufio.NewWriter(os.Stdout)
	for i := 1; i <= 1000000; i++ {
		writer.WriteString(strconv.Itoa(i) + "\n")
	}
	writer.Flush()

}
