package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 2)
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	done := make(chan bool, 1)
	go worker(done)
	<-done

}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
