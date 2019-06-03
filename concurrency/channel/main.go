package main

import (
	"fmt"
	"sync"
)

func greeting(myChannel chan string) {
	myChannel <- "hello"
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	// channel1 := make(chan string)
	// channel2 := make(chan string)
	// go abc(channel1)
	// go def(channel2)
	// fmt.Print(<-channel1)
	// fmt.Print(<-channel2)
	// fmt.Print(<-channel1)
	// fmt.Print(<-channel2)
	// fmt.Print(<-channel1)
	// fmt.Print(<-channel2)
	// fmt.Println()
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()

}
