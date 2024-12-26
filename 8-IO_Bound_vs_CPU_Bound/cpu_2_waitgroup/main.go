package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(16)

	var wg sync.WaitGroup

	start := time.Now()

	// Start CPU-bound tasks in goroutines
	tasks := []string{"Task A", "Task B", "Task C"}
	for _, task := range tasks {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			count(name)
		}(task)
	}

	// Wait for all tasks to complete
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Processes took: %s\n", elapsed)
}

func count(name string) {
	fmt.Printf("%s is starting\n", name)
	for i := 1; i < 10_000_000_000; i++ {
	}
	fmt.Printf("%s is done\n", name)
}
