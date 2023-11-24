package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	text = strings.Replace(text, ",", "", -1)
	text = strings.Replace(text, ".", "", -1)
	words := strings.Split(text, " ")
	result := make([]int, len(words))
	for i, w := range words {
		result[i] = len(w)
	}
	fmt.Println(result)
}
