package main

import (
	"fmt"
	"strings"
)

func main() {
	// name := "Chapman, Luke"
	// score := 90

	type score struct {
		name  string
		score int
	}

	// students := []string{
	// 	"Chapman, Luke",
	// 	"Smith, John",
	// 	"Kettle, James", // notice the trailing comma at the end of the array, for strings we need this
	// }
	scores := []score{
		{name: "Chapman, Luke", score: 87},
		{name: "Smith, John", score: 85},
		{name: "Kettle, James", score: 99},
	}

	// score := []int{88, 83, 79}
	// name, score := "Chapman, Luke", 90

	fmt.Println("Student Scores")
	fmt.Println(strings.Repeat("-", 17))
	fmt.Println(scores[0].name, scores[0].score)
	fmt.Println(scores[1].name, scores[1].score)
	fmt.Println(scores[2].name, scores[2].score)
}
