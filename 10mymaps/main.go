package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	langusges := make(map[string]string)

	langusges["JS"] = "Javascript"
	langusges["RB"] = "Ruby"
	langusges["PY"] = "Python"

	fmt.Println(langusges)
	fmt.Println("JS shorts for: ", langusges["JS"])

	delete(langusges, "RB")

	fmt.Println("Deleted: ", langusges)

	// loops are interesting in golang
	for key, value := range langusges {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	for _, value := range langusges {
		fmt.Printf("For key No, value is %v\n", value)
	}
}
