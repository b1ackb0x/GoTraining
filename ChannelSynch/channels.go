package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
}

func main() {

	// Start a worker goroutine
	done := make(chan bool, 1) // buffered channel
	go worker(done)

	// Block until we receive a notification from the
	// worker on the channel.
	<-done
}