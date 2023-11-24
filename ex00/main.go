package main

import (
	"fmt"
)

func reverse(s string) string {
	ss := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		ss[i], ss[len(s)-1-i] = ss[len(s)-1-i], ss[i]
	}
	return string(ss)
}

func main() {
	fmt.Println(reverse("stressed"))
}
