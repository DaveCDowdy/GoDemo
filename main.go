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
	

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(scores[0].name, scores[0].score)
	fmt.Println(scores[1].name, scores[1].score)
	fmt.Println(scores[2].name, scores[2].score)
}
