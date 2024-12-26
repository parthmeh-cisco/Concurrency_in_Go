// // This is concurrency / goroutines implemented with channels
package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func main() {
// 	fmt.Println(runtime.GOMAXPROCS(0))
// 	runtime.GOMAXPROCS(16) // Extra processors help up to number of goroutines
// 	c := make(chan string)
// 	startTime := time.Now()
// 	go counta(c)
// 	go countb(c)
// 	go countc(c)
// 	go countd(c)
// 	go counte(c)
// 	go countf(c)
// 	go countg(c)
// 	go counth(c)

// 	for i := 0; i < 8; i++ {
// 		fmt.Println(<-c)
// 	}
// 	elapsed := time.Since(startTime)
// 	fmt.Printf("Processes took: %s\n", elapsed)
// }

// func counta(c chan string) {
// 	fmt.Println("AAAA is starting")
// 	for I := 1; I < 10_000_000_000; I++ {
// 	}

// 	c <- "AAAA is done"
// }

// func countb(c chan string) {
// 	fmt.Println("BBBB is starting")
// 	for I := 1; I < 10_000_000_000; I++ {
// 	}

// 	c <- "BBBB is done"
// }

// func countc(c chan string) {
// 	fmt.Println("CCCC is starting")
// 	for I := 1; I < 10_000_000_000; I++ {
// 	}

// 	c <- "CCCC is done"
// }

// func countd(c chan string) {
// 	fmt.Println("DDDD is starting")
// 	for I := 1; I < 10_000_000_000; I++ {
// 	}

// 	c <- "DDDD is done"
// }

// func counte(c chan string) {
// 	fmt.Println("EEEE is starting")
// 	for I := 1; I < 10_000_000_000; I++ {
// 	}

// 	c <- "EEEE is done"
// }

// func countf(c chan string) {
// 	fmt.Println("FFFF is starting")
// 	for I := 1; I < 10_000_000_000; I++ {
// 	}

// 	c <- "FFFF is done"
// }

// func countg(c chan string) {
// 	fmt.Println("GGGG is starting")
// 	for I := 1; I < 10_000_000_000; I++ {
// 	}

// 	c <- "GGGG is done"
// }

// func counth(c chan string) {
// 	fmt.Println("HHHH is starting")
// 	for I := 1; I < 10_000_000_000; I++ {
// 	}

// 	c <- "HHHH is done"
// }
