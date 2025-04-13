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

	for shouldContinue {

		fmt.Println("1) Enter a score")
		fmt.Println("2) Print a report")
		fmt.Println("q) Quit")
		fmt.Println()
		fmt.Println("Enter an option")

		var option string
		fmt.Scanln(&option)

		switch option {
		case "1":
			scores = append(scores, addStudent())
		case "2":
			printReport(scores)
		case "q":
			shouldContinue = false
		}
	}
}

func printReport(scores []score) {
	fmt.Println("Student scores")
			fmt.Println(strings.Repeat("-", 14))
			for _, s := range scores {
				fmt.Println(s.name, s.score)
			}
}

func addStudent() score {
	fmt.Println("Enter a student name and score")
	var name, rawScore string
	fmt.Scanln(&name, &rawScore)
	scoreInt, _ := strconv.Atoi(rawScore)
	return score{name: name, score: scoreInt}
}
