package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		channel <- 7
		channel <- 5
		channel <- 4
		close(channel)
	}()

	avg(channel)
}

func avg(channel chan int) {
	var sum int
	var count int

	for value := range channel {
		sum += value
		count++
		fmt.Println("Rata2 : ", sum/count)
	}
}
