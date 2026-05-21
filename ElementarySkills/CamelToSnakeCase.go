package main

import (
	"fmt"
	"strings"
)

// Instructions
// Write a function that converts a string from camelCase to snake_case.
// -If the string is empty, return an empty string.
// -If the string is not camelCase, return the string unchanged.
// -If the string is camelCase, return the snake_case version of the string.

// For this exercise you need to know that
// camelCase has two different writing alternatives that will be accepted:
// -lowerCamelCase
// -UpperCamelCase
// -Rules for writing in camelCase:

// The word does not end on a capitalized letter (CamelCasE).
// No two capitalized letters shall follow directly each other (CamelCAse).
// Numbers or punctuation are not allowed in the word anywhere (camelCase1).

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	result := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		// Check if uppercase
		if ch >= 'A' && ch <= 'Z' {

			// Invalid if uppercase is last character
			if i == len(s)-1 {
				return s
			}
			// Invalid if next character is also uppercase
			if i+1 < len(s) && s[i+1] >= 'A' && s[i+1] <= 'Z' {
				return s
			}

			// Add underscore except at beginning
			if i != 0 {
				result += "_"
			}

			// Convert to lowercase
			//result += strings.ToLower(string(ch))

		} else if ch >= 'a' && ch <= 'z' {

			result += string(ch)

		} else {
			// Invalid character
			return s
		}
	}

	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

//Output
// $ go run .
// Hello_World
// hello_World
// camel_Case
// CAMELtoSnackCASE
// camel_To_Snake_Case
// hey2
