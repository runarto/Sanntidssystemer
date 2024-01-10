package main

import (
	"fmt"
	"time"
)

func producer(buffer chan int) {

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		buffer <- i
		// TODO: push real value to buffer
	}
	close(buffer)

}

func consumer(buffer chan int) {

	time.Sleep(1 * time.Second)
	for {
		i := <-buffer //TODO: get real value from buffer
		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}
	close(buffer)

}

func main() {

	// TODO: make a bounded buffer
	bufferSize := 5
	buffer := make(chan int, bufferSize)

	go consumer(buffer)
	go producer(buffer)

	select {}
}
