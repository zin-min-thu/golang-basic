package main

import "fmt"

func main() {
	fmt.Println("Function in golang")
	greeter()

	result := adder(3, 4)
	fmt.Println("Result is: ", result)

	proResult, message := proAdder(2, 3, 4, 5)
	fmt.Println("Pro Result is: ", proResult)
	fmt.Println("Pro Message is: ", message)
}

func greeter() {
	fmt.Println("Hello greeting")
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

// array vairables in function
func proAdder(values ...int) (int, string) {

	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Here Pro adder message"
}
