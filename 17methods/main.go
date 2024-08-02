package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang, no super or parent
	hitesh := User{"Hitesh", "hitesh@gmail.com", true, 15}

	fmt.Println(hitesh)
	fmt.Printf("Detail: %+v\n", hitesh)
	fmt.Printf("Name: %v and email: %v\n", hitesh.Name, hitesh.Email)

	hitesh.GetStatus()

	hitesh.NewMail()

	fmt.Printf("Name: %v and email: %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "newemail@gmail.com"

	fmt.Println("New Email is: ", u.Email)
}
