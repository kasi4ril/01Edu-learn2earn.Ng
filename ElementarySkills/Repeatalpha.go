package main

import "fmt"

// Instructions
// Write a function called RepeatAlpha that takes a string and displays it
// repeating each alphabetical character as many times as its alphabetical index.
// 'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...

func RepeatAlpha(s string) string {

	result := ""

	for _, ch := range s {

		// searching for lowercase letters
		if ch >= 'a' && ch <= 'z' {

			count := int(ch-'a') + 1

			for j := 0; j < count; j++ {
				result += string(ch)
			}

		// searching for uppercase letters
		} else if ch >= 'A' && ch <= 'Z' {

			count := int(ch-'A') + 1

			for j := 0; j < count; j++ {
				result += string(ch)
			}

		// for non-letters
		} else {
			result += string(ch)
		}
	}

	return result
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}

//Output
// $ go run . | cat -e
// abbccc$
// CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
// $
// abbacccaddddabba 01!$
// $
