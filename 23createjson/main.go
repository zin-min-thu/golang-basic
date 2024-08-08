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

	// EncodeJson()
	DecodJson()
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

func DecodJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJs Bootcamp",
			"Price": 222,
			"website": "Google.com",
			"tags": ["Web-dev","js"]
        }
	`)

	var locCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Valid Json Format")
		json.Unmarshal(jsonDataFromWeb, &locCourse)

		fmt.Printf("%#v\n", locCourse)
	} else {
		fmt.Println("Invalid json fromat")
	}

	// key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
