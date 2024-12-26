package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(16)

	links := []string{
		"http://hashnode.com",
		"http://dev.to",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://medium.com",
		"http://github.com",
		"http://techcrunch.com",
		"http://techrepublic.com",
	}

	var wg sync.WaitGroup

	start := time.Now()

	for _, link := range links {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			checkLink(l)
		}(link)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Processes took: %s\n", elapsed)
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not responding!")
		return
	}
	fmt.Println(link, "is LIVE!")
}
