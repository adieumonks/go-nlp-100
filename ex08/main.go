package main

import (
	"fmt"
	"unicode"
)

func main() {
	encoded := cipher("I am an NLPer")
	decoded := cipher(encoded)
	fmt.Printf("encoded: %s\n", encoded)
	fmt.Printf("decoded: %s\n", decoded)
}

func cipher(text string) string {
	result := []rune(text)
	for i, r := range text {
		if unicode.IsLower(r) {
			result[i] = 219 - r
		}
	}
	return string(result)
}
