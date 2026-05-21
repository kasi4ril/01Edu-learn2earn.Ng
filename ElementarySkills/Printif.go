package main

import "fmt"

//Instructions
//Write a function that takes a string as an argument and returns the letter G followed by a newline \n
// if the argument length is more or equal than 3, otherwise returns Invalid Input followed by a newline \n.
//If it's an empty string return G followed by a newline \n....

func PrintIf(str string) string {

	if len(str) >= 3 {
		return "G\n"
	}
	if str == ""{
		return "G\n"
	}
	return "Invalid input\n"
}

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}

//Output
//$ go run . | cat -e
//G$
//G$
//G$
//Invalid Input$
