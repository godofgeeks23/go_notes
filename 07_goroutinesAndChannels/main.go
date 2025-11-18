package main

import (
	"fmt"
	"time"
)

func doSomething() {
	fmt.Println("entered into the doSomething() func")
	fmt.Println("line 3 from something function")
}

func slowTask() int {
	fmt.Println("entered into the slowTask() func")
	time.Sleep(2 * time.Second)
	return 2
}

func main() {
	fmt.Println("line 1 from main")

	// any function can run concurrently as a goroutine (lightweight thread). so in a way go = async ()
	go doSomething()
	fmt.Println("line 2 from main")

	// to sync a result of a goroutine, channels are used
	ch := make(chan int)
	go func() {
		ch <- slowTask()
	}()

	result := <-ch
	fmt.Println(result)

}
