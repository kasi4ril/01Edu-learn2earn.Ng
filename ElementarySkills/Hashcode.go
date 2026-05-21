package main

import "fmt"

// Instructions
// Write a function called HashCode() that takes a string as
//  an argument and returns a new hashed string.
// The hash equation is computed as follows:
// (ASCII of current character + size of the string) % 127,
// ensuring the result falls within the ASCII range of 0 to 127.

// If the resulting character is unprintable add 33 to it.

func HashCode(dec string) string {

	size_of_string := len(dec)
	result := ""
	for _, ch := range dec {

		hash := (int(ch) + size_of_string) % 127
		if hash < 32 {
			hash += 33
		}

		result += string(rune(hash))

	}
	return result
}
func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

//Output
// $ go run .
// B
// CD
// EDF
// Spwwz+bz}wo
