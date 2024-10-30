package main

import (
	"fmt"
)

func main() {
	animalChannel := make(chan string)
	fruitChannel := make(chan string)

	go func() {
		animalChannel <- "cat"
		animalChannel <- "dog"
		animalChannel <- "fish"
		animalChannel <- "chicken"
		animalChannel <- "bird"
		close(animalChannel)
	}()

	go func() {
		fruitChannel <- "banana"
		fruitChannel <- "grapes"
		fruitChannel <- "apple"
		fruitChannel <- "orange"
		fruitChannel <- "durian"
		close(fruitChannel)
	}()

	fmt.Println("Process")

	var animalStatus bool
	var fruitStatus bool

	for {
		select {
		case data, status := <-animalChannel:
			animalStatus = status
			if animalStatus {
				fmt.Println("Animal : " + data)
			}
		case data, status := <-fruitChannel:
			fruitStatus = status
			if fruitStatus {
				fmt.Println("fruit : " + data)
			}
			if !animalStatus && !fruitStatus {
				break
			}
		}
	}

	fmt.Println("Done")
}
