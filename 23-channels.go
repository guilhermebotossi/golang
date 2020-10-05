package main

import "fmt"
import "time"

func main() {

	messages := make(chan string)

	go func() { 
		fmt.Println("Antes do channel")
		time.Sleep(10)
		messages <- "ping" 
		time.Sleep(10)
		fmt.Println("depois do sleep")}()


	fmt.Println("antes")
    msg := <-messages
	fmt.Println(msg)
	fmt.Println("depois")
}