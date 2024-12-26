package main

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func main() {

	start := time.Now()
	go doSomething()
	go doSomethingElse()

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("\n\nI guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took: %s\n", elapsed)
}

func doSomething() {
	time.Sleep(2 * time.Second)
	fmt.Println("\nI've done something")
	ch <- "do Something finished"
}

func doSomethingElse() {
	time.Sleep(2 * time.Second)
	fmt.Println("I've done something else")
	ch <- "do SomethingElse finished"
}
