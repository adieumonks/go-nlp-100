package main

import "fmt"

func main() {
	text := "パタトクカシーー"
	runes := []rune(text)
	result := make([]rune, len(runes)/2)
	for i := 0; i < len(runes); i += 2 {
		result[i/2] = runes[i]
	}
	fmt.Println(string(result))
}
