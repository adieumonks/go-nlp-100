package main

import "fmt"

func main() {
	runes1 := []rune("パトカー")
	runes2 := []rune("タクシー")
	result := make([]rune, len(runes1)+len(runes2))
	for i := 0; i < len(runes1); i++ {
		result[2*i] = runes1[i]
		result[2*i+1] = runes2[i]
	}
	fmt.Println(string(result))
}
