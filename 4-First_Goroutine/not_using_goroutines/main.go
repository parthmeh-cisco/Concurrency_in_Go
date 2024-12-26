package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	go doSomething()
	go doSomethingElse()

	time.Sleep(5 * time.Second)

	fmt.Println("\n\nI guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took: %s\n", elapsed)
}

func doSomething() {
	time.Sleep(2 * time.Second)
	fmt.Println("\nI've done something")
}

func doSomethingElse() {
	time.Sleep(2 * time.Second)
	fmt.Println("I've done something else")
}
