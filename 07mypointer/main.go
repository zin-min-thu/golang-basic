package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)  // momery address
	fmt.Println("Value of actual pointer is ", *ptr) // check pointer value

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)
}
