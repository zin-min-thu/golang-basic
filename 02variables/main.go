package main

import "fmt"

const LoginToken string = "testingladfskla"

func main() {
	var username string = "Ko Zin"

	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true

	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255

	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.99999

	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int

	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "www.google.com"

	fmt.Println(website)
	fmt.Printf("variable is of type: %T \n", website)

	// no var style

	numberOfUser := 30000
	fmt.Println(numberOfUser)

	// public variable
	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)
}
