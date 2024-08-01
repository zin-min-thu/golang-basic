package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice golang")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println("Fruit list ", fruitList)

	fruitList = append(fruitList[1:4])

	fmt.Println("Fruit list second ", fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 444
	highScore[2] = 555
	highScore[3] = 666
	// highScore[4] = 238 // can not insert

	highScore = append(highScore, 238, 239, 240) // relocate memory can insert new

	fmt.Println(highScore)

	// sort
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	// remove value from slice based on index
	var courses = []stirng{"ReactJs", "Javascirpt", "Swift", "Python", "Ruby"}
	fmt.Println(courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
