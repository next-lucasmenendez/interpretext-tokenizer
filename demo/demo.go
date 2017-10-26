package main

import (
	"fmt"
	"github.com/lucasmenendez/gotokenizer"
)

func main() {
	var input string = "Korea electronics giant LG continues to show the world that making a profit, or even breaking even, from selling smartphones is no easy thing."

	var sentences [][]string = gotokenizer.Sentences(input)
	for _, s := range sentences {
		fmt.Print("[")
		for _, t := range s {
			fmt.Printf("%s, ", t)
		}
		fmt.Print("]\n")
	}
}
