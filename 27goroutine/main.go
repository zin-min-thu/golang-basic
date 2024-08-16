package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointer

var mut sync.Mutex // pointer

func main() {
	// go greeter("Hello")
	// greeter("World")

	websitelist := []string{
		"https://google.com",
		"https://facebook.com",
		"https://laravel.com",
		"https://github.com",
		"https://tommy.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {

	defer wg.Done()

	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Println(err)
	}

	mut.Lock()

	signals = append(signals, endpoint)

	mut.Unlock()

	fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
}
