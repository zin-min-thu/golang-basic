package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zin-min-thu/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")

	r := router.Router()

	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is listening at prot 4000...")
}
