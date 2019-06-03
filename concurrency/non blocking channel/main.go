package main

import (
	"fmt"
	"math/rand"
	"time"
)

func cakeMaker(kind string, number int, to chan<- string) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < number; i++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Microsecond)
		to <- kind
	}
	close(to)
	fmt.Println("close channel>>>>>")
}

func main() {
	chocolate := make(chan string, 8)
	revelvet := make(chan string, 8)

	go cakeMaker("chocolate", 4, chocolate)
	go cakeMaker("red velvet", 40, revelvet)

	moreChocolate := true
	moreRevelvet := true
	var cake string

	for moreChocolate || moreRevelvet {
		select {
		case cake, moreChocolate = <-chocolate:
			if moreChocolate {
				fmt.Printf("Get a cake from the first factory: %s\n", cake)
			}

		case cake, moreRevelvet = <-revelvet:
			if moreRevelvet {
				fmt.Printf("Get a cake from the first factory: %s\n", cake)
			}

		}
	}
}
