package main

import (
	"fmt"
	"strings"
)

func main() {
	// name := "Chapman, Luke"
	// score := 90

	name, score := "Chapman, Luke", 90

	fmt.Println("Student Scores")
	fmt.Println(strings.Repeat("-", 17))
	fmt.Println(name, score)
}
