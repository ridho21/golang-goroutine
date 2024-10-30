package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	channel := make(chan string)
	defer close(channel)

	wg.Add(2)
	// go gateway(channel, "Hello, Ridho")
	go sendMessage(channel, "Whatsap")
	go getMessage(channel)

	// go func() {
	// 	channel <- "Hello World!"
	// }()

	// words := <-channel
	// fmt.Println(<-channel)
	wg.Wait()
	fmt.Println("DONE")
}

func gateway(channel chan string, words string) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
	channel <- words
}

func sendMessage(channel chan<- string, message string) {
	defer wg.Done()
	channel <- message
}

func getMessage(channel <-chan string) {
	defer wg.Done()
	fmt.Println(<-channel)
}
