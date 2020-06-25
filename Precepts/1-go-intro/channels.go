package main

import "fmt"

// Launch n goroutines, each printing a number
// Note how the numbers are not printed in order
func goroutines() {
	for i := 0; i < 10; i++ {
		// Print the number asynchronously
		go fmt.Printf("Printing %v in a goroutine\n", i)
	}
	// At this point the numbers may not have been printed yet
	fmt.Println("Launched the goroutines")
}

// Channels are a way to pass messages across goroutines
func channels() {
	ch := make(chan int)
	// Launch a goroutine using an anonymous function
	go func() {
		i := 1
		for {
			// This line blocks until someone
			// consumes from the channel
			ch <- i * i
			i++
		}
	}()
	// Extract first 10 squared numbers from the channel
	for i := 0; i < 10; i++ {
		// This line blocks until someone sends into the channel
		fmt.Printf("The next squared number is %v\n", <-ch)
	}
}

// Buffered channels are like channels except:
//   1. Sending only blocks when the channel is full
//   2. Receiving only blocks when the channel is empty
func bufferedChannels() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	// Buffer is now full; sending any new messages will block
	// Instead let's just consume from the channel
	for i := 0; i < 3; i++ {
		fmt.Printf("Consuming %v from channel\n", <-ch)
	}
	// Buffer is now empty; consuming from channel will block
}
