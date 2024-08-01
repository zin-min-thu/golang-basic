package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang, no super or parent
	hitesh := User{"Hitesh", "hitesh@gmail.com", true, 15}

	fmt.Println(hitesh)
	fmt.Printf("Detail: %+v\n", hitesh)
	fmt.Printf("Name: %v and email: %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
