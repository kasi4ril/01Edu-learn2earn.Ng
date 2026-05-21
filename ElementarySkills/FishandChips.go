package main

import "fmt"

// Instructions
// Write a function called FishAndChips() that takes an int and returns a string.

// If the number is divisible by 2, print fish.
// If the number is divisible by 3, print chips.
// If the number is divisible by 2 and 3, print fish and chips.
// If the number is negative return error: number is negative.
// If the number is non divisible by 2 or 3 return error: non divisible.

func FishAndChips(n int) string {

	result := ""
	if n < 0 {
		result = "number is negative"
	}
	//you first check for both numbers so that the logic falls true for single case
	if n%2 == 0 && n%3 == 0 {
		result = "fish and chips"
	} else if n%2 == 0 {
		result = "fish"
	} else if n%3 == 0 {
		result = "chips"
	}

	return result
}
func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}

//Output
// $ go run . | cat -e
// fish$
// chips$
// fish and chips$
