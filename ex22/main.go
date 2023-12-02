package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

type Article struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func main() {
	f, err := os.Open("../data/jawiki-country.json")
	if err != nil {
		fmt.Printf("failed to open file: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1000000), 1000000)
	var article Article
	for sc.Scan() {
		if err := json.Unmarshal(sc.Bytes(), &article); err != nil {
			fmt.Printf("failed to unmarshal json: %v", err)
			os.Exit(1)
		}
		if article.Title == "イギリス" {
			break
		}
	}

	pattern := `\[\[Category:(.*)\]\]`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("failed to compile regexp: %v", err)
		os.Exit(1)
	}

	matches := re.FindAllStringSubmatch(article.Text, -1)
	for _, match := range matches {
		fmt.Println(match[1])
	}
}
