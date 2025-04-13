package main

import (
	"fmt"
	"strings"
)

func main() {

	type score struct {
		name string
		score int
	}

	scores := []score{
		{name: "Dent, Arthur", score: 87},
		{name: "McMillan, Tricia", score: 96},
		{name: "Prefect, Ford", score: 64},
	}
	
	fmt.Println("Select a score to print (1-3)")
	var option string
	fmt.Scanln(&option)

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 14))
	var index int
	switch option{
	case "1":
		index = 0
	case "2":
		index = 1
	case "3":
		index = 2
	default:
		fmt.Println("Invalid option defaulting to 1")
		index = 0
	}
	
	fmt.Println(scores[index].name, scores[index].score)
}
