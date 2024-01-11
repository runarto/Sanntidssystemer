// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

type Action int

const (
	Increment Action = 0 // Enum for increment action
	Decrement Action = 1 // Enum for decrement action
	Read      Action = 2 // Enum for read action
)

type Message struct {
	action   Action
	response chan int // Channel for sending back the read value
}

var (
	actionChan = make(chan Message) // Channel for sending actions to the numberServer
	doneChan   = make(chan bool)    // Channel to signal completion of incrementing/decrementing
)

func numberServer() {
	var i int // The number being incremented or decremented
	for {
		msg := <-actionChan // Wait and receive a message from the channel
		switch msg.action {
		case Increment:
			i++ // Increment action
		case Decrement:
			i-- // Decrement action
		case Read:
			msg.response <- i // Send the current value of i back through the channel
		}
	}
}

func incrementing() {
	for j := 0; j < 1000000; j++ {
		actionChan <- Message{action: Increment} // Send increment action 1 million times
	}
	doneChan <- true // Signal completion of incrementing
}

func decrementing() {
	for j := 0; j < 999999; j++ {
		actionChan <- Message{action: Decrement} // Send decrement action 999999 times
	}
	doneChan <- true // Signal completion of decrementing
}

func main() {
	// GOMAXPROCS sets the maximum number of CPUs that can execute simultaneously
	runtime.GOMAXPROCS(2)

	// Starting the numberServer and incrementing/decrementing goroutines
	go numberServer()
	go incrementing()
	go decrementing()

	// Wait for both incrementing and decrementing goroutines to finish
	<-doneChan
	<-doneChan

	responseChan := make(chan int)                              // Channel to store value
	actionChan <- Message{action: Read, response: responseChan} // Request to read the current value of i
	i := <-responseChan                                         // Receive the value of i

	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
