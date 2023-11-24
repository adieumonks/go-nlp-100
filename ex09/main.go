package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	text := "I couldn’t believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	words := strings.Split(text, " ")

	result := make([]string, len(words))
	for i, w := range words {
		if len(w) > 4 {
			result[i] = shuffle(w)
		} else {
			result[i] = w
		}
	}
	fmt.Println(strings.Join(result, " "))
}

func shuffle(word string) string {
	result := []rune(word)
	for i := 1; i < len(result)-1; i++ {
		j := 1 + rand.Intn(len(result)-2)
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
