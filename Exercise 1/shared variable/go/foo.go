// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

type Action int

const (
	Increment Action = iota
	Decrement
	Read
)

type Message struct {
	action   Action
	response chan int // Used for Read action
}

var (
	actionChan = make(chan Message)
	doneChan   = make(chan bool)
)

func numberServer() {
	var i int
	for {
		msg := <-actionChan
		switch msg.action {
		case Increment:
			i++
		case Decrement:
			i--
		case Read:
			msg.response <- i
		}
	}
}

func incrementing() {
	for j := 0; j < 1000000; j++ {
		actionChan <- Message{action: Increment}
	}
	doneChan <- true
}

func decrementing() {
	for j := 0; j < 999999; j++ {
		actionChan <- Message{action: Decrement}
	}
	doneChan <- true
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(1)

	// TODO: Spawn both functions as goroutines

	go numberServer()
	go incrementing()
	go decrementing()

	// Wait for both goroutines to finish
	<-doneChan
	<-doneChan

	responseChan := make(chan int)
	actionChan <- Message{action: Read, response: responseChan}
	i := <-responseChan

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
