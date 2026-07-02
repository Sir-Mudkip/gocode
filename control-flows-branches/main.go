package main

import (
	"fmt"
	"strings"
)

func main() {
	type score struct {
		name  string
		score int
	}
	scores := []score{
		{name: "Chapman, Luke", score: 87},
		{name: "Smith, John", score: 85},
		{name: "Kettle, James", score: 99},
	}

	fmt.Println("Select score to print (1 - 3):")
	var option string   // delcare our variable
	fmt.Scanln(&option) // takes user input from CLI, points to memory location of "option"

	fmt.Println("Student Scores")
	fmt.Println(strings.Repeat("-", 17))
	var index int
	//if option == "1" { // use strconv package if using on something like a website
	//	index = 0
	//} else if option == "2" {
	//	index = 1
	//} else if option == "3" {
	//	index = 2
	//} else {
	//	fmt.Println("Unknown Option, defaulting to 1")
	//	index = 0
	//}
	switch option {
	case "1":
		index = 0
	case "2":
		index = 1
	case "3":
		index = 2
	default:
		fmt.Println("Unknown value, defaulting to 1")
		index = 0
	}
	fmt.Println(scores[index].name, scores[index].score)
}
