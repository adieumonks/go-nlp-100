package main

import "fmt"

func ngram(text string, n int) []string {
	result := make([]string, len(text)-n+1)
	for i := 0; i < len(text)-n+1; i++ {
		result[i] = text[i : i+n]
	}
	return result
}

func main() {
	fmt.Printf("%v", ngram("I am an NLPer", 2))
}
