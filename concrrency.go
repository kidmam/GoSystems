package main

//https://medium.com/fstack/concrrency-in-go-goroutines-part-1-ea94606d90d

import (
	"fmt"
	"sync"
)

func printHelloWorld() {
	fmt.Println("Hello World")
}

/*func main() {
	go printHelloWorld()
	go func() {
		fmt.Println("Hello World")
	}()
}*/

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	//fork point
	go func() {
		defer wg.Done()
		fmt.Println("second goroutine")
	}()

	//join point
	wg.Wait()
	fmt.Println("main goroutine")
}
