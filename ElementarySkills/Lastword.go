package main

import "fmt"

// Instructions
// Write a function LastWord that takes a string and returns its last word followed by a \n.
// A word is a section of string delimited by spaces or by the start/end of the string.

// LastWord returns the last word in a string
func LastWord(s string) string {

	end := len(s) - 1
	// Skip trailing spaces
	for end >= 0 && s[end] == ' ' {
		end--
	}
	// If string is empty or only spaces
	if end < 0 {
		return ""
	}

	start := end
	// Move backwards until space is found
	for start >= 0 && s[start] != ' ' {
		start--
	}

	return s[start+1:end+1] + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

//Output
// $ go run . | cat -e
// not$
// lorem,ipsum$
// $
// $
