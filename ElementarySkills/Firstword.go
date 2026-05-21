package main

import "fmt"

// Instructions
// Write a function that takes a string and return a string
// containing its first word, followed by a newline ('\n').

// A word is a sequence of characters delimited by spaces or by
// the start/end of the argument.
// Expected Function

func FirstWord(s string) string {

	//go through the string and stop immediately after meeting a space
	for i, _ := range s {
		if s[i] == ' '{
			return s[:i]+"\n"
		}
	}
	return s[:]+"\n"
}
func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}
//Output
// $ go run .
// hello

// hello
// $
