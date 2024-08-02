package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://p260.visibleone.io"

func main() {
	fmt.Println("MY Web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", response)

	// defer last in first out
	defer response.Body.Close() // caller's responsibility to close the connectivity

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
