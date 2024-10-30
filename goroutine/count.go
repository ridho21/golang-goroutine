package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go repeat("cats", 500, &wg)
	go repeat("dogs", 500, &wg)
	wg.Wait()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done")
}

func repeat(word string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println(i, word)
		time.Sleep(time.Millisecond * delay)
	}
}
