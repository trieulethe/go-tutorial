package main

import (
	"fmt"
	"net/http"
	"sync"
)

func webGetWorker(in chan string, wg *sync.WaitGroup) {
	fmt.Println("chan in", in)
	for {
		url := <-in
		res, err := http.Get(url)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("GET %s: %d\n", url, res.StatusCode)
		}

		wg.Done()
	}
}

func main() {
	work := make(chan string, 1024)
	var wg sync.WaitGroup
	numWorker := 1400
	for i := 0; i < numWorker; i++ {
		go webGetWorker(work, &wg)
	}

	urls := [4]string{"http://packpub.com", "http://reddit.com", "http://twiter.com", "http://example.com"}
	for i := 0; i < 100; i++ {
		for _, url := range urls {
			wg.Add(1)
			work <- url
		}
	}

	wg.Wait()
}
