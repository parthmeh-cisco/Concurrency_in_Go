package main

import (
	"fmt"
	"net/http"
	"runtime"
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

	c := make(chan string, len(links))

	start := time.Now()

	for _, link := range links {
		go checkLink(link, c)
	}

	for len(c) < len(links) { // Program ends when everything is in BUFFERS channel.
	}

	// for range links { // This also works beautifully, using blocking code
	// 	fmt.Println("channel message:",<-c)
	// }

	elapsed := time.Since(start)
	fmt.Printf("Processes took: %s\n", elapsed)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not responding!")
		c <- link + " is not responding!"
		return
	}

	c <- link + " is LIVE!"
}
