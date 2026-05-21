package main

import (
	"fmt"
)

//Instructions
//Write a function that takes a string and returns:
//"G\n" if the string's length is less than 3 (including empty string).
//"Invalid Input\n" otherwise.

func PrintIfNot(str string) string {

	if len(str) < 3 || str == "" {
		return "G\n"
	}
	return "Invalid Input\n"
}
func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}
// its output:

// $ go run . | cat -e
// Invalid Input$
// Invalid Input$
// G$
// G$