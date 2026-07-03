package main

import (
	"fmt"
	"strconv"
	"strings"
)

type score struct {
	name  string
	score int
}

func main() {

	scores := []score{}
	shouldContinue := true

	for shouldContinue { // declares our CLI, while true it loops the following logic

		fmt.Println("1) Enter a Score")
		fmt.Println("2) Print Report")
		fmt.Println("q) Quit")
		fmt.Println()
		fmt.Println("Please select an option")

		var option string   // declare our options for the CLI
		fmt.Scanln(&option) // Take option as a user-defined input

		switch option {
		case "1":
			scores = append(scores, addStudent()) // redefine scores, by appending it's struct with addStudent
		case "2":
			printReport(scores)
		case "q":
			shouldContinue = false // turns the condition to false, then exits the loop, closing the logic
		}
	}
}

func addStudent() score {
	fmt.Println("Enter student name and score")
	var name, rawScore string    // declare names and raw scores
	fmt.Scanln(&name, &rawScore) // take above vars as user input
	s, _ := strconv.Atoi(rawScore)
	return score{name: name, score: s}
}

func printReport(scores []score) {
	fmt.Println("Student Scores")
	fmt.Println(strings.Repeat("=", 20))
	for _, s := range scores {
		fmt.Println(s.name, s.score)
	}
}
