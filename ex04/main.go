package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	text := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	text = strings.Replace(text, ".", "", -1)
	words := strings.Split(text, " ")
	result := make(map[string]int)
	targetNums := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}
	for i, w := range words {
		if slices.Contains(targetNums, i+1) {
			result[string(w[0])] = i
		} else {
			result[w[:2]] = i
		}
	}
	fmt.Println(result)
}
