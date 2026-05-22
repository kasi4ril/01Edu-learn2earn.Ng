package main

import (
	"fmt"
	"os"
)

// Instructions
// Write a program that takes a string and displays it with exactly three spaces
// between each word, with no spaces nor tabs at neither the beginning nor the end.
// The string will be followed by a newline ('\n').
// A word, in this exercise, is a sequence of visible characters.
// If the number of arguments is not 1, or if there are no word, the program displays nothing.

func main() {

	// Must have exactly 1 argument
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]
	i := 0
	wordFound := false
	// Skip leading spaces/tabs
	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}

	for i < len(s) {

		// Print word characters
		if s[i] != ' ' && s[i] != '\t' {

			wordFound = true
			fmt.Print(string(s[i]))
			i++
		} else {
			// Skip all spaces/tabs
			for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
				i++
			}
			// Print exactly 3 spaces if another word exists
			if i < len(s) {
				fmt.Print("   ")
			}
		}
	}
	// Print newline only if at least one word exists
	if wordFound {
		fmt.Println()
	}
}

// Usage
// $ go run . "you   see   it's   easy   to   display   the   same   thing" | cat -e
// you   see   it's   easy   to   display   the   same   thing$
// $ go run . "   only  it's harder   " | cat -e
// only   it's   harder$
// $ go run . " how funny it is" "did you  hear, Mathilde ?" | cat -e
// $ go run .
// $
/// ALSO USED FOR TEXT NORMALIZATION IN SEARCH ENGINES, DATA PROCESSING ETC
