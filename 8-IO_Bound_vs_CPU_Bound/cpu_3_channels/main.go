package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(16)

	start := time.Now()

	// Create a channel to communicate completion of tasks
	done := make(chan bool)

	// Start CPU-bound tasks in goroutines
	go count("Task A", done)
	go count("Task B", done)
	go count("Task C", done)

	// Wait for all tasks to complete
	for i := 0; i < 3; i++ {
		<-done
	}

	elapsed := time.Since(start)
	fmt.Printf("Processes took: %s\n", elapsed)
}

func count(name string, done chan bool) {
	fmt.Printf("%s is starting\n", name)
	for i := 1; i < 10_000_000_000; i++ {
	}
	fmt.Printf("%s is done\n", name)
	done <- true
}
