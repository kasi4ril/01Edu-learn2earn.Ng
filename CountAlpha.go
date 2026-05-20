package main

import (
	"fmt"
)
//Instructions
//Write a function CountAlpha() that takes a string as an argument and returns 
//the number of alphabetic characters in the string.

func CountAlpha(s string) int {

	count := 0
	
	for _, value := range s {
		if value == ' ' {
			continue
		}
		if value <= '0' || value <= '9' {
			continue
		}
		if value <= 'a' || value <= 'z' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}

//output
// $ go run .
// 10
// 5
// 5
