package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 3)

	go foo(messages)


	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func foo(messages chan string) {
	messages <- "buffered"
	time.Sleep(5 * time.Second)
	messages <- "channel"
}
