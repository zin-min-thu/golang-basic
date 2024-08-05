package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Learn JSON in golang")

	EncodeJson()
}

func EncodeJson() {
	locCourses := []course{
		{"ReactJs Bootcamp", 222, "Google.com", "55sdaflk", []string{"Web-dev", "js"}},
		{"Java Bootcamp", 333, "Google.com", "12dddd", []string{"Web-dev", "js"}},
		{"Python Bootcamp", 444, "Google.com", "44cccc", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(locCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}
