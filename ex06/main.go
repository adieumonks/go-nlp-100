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
	X := ngram("paraparaparadise", 2)
	Y := ngram("paragraph", 2)
	fmt.Printf("X: %v\n", X)
	fmt.Printf("Y: %v\n", Y)
	fmt.Printf("Union: %v\n", union(X, Y))
	fmt.Printf("Intersection: %v\n", intersection(X, Y))
	fmt.Printf("Difference: %v\n", difference(X, Y))
	fmt.Printf("se in X: %v, se in Y: %v\n", in(X, "se"), in(Y, "se"))
}

func union(X, Y []string) []string {
	resultMap := make(map[string]bool)
	for _, x := range X {
		resultMap[x] = true
	}
	for _, y := range Y {
		resultMap[y] = true
	}
	return keys(resultMap)
}

func intersection(X, Y []string) []string {
	result := []string{}
	resultMap := make(map[string]bool)
	for _, x := range X {
		resultMap[x] = true
	}
	for _, y := range Y {
		if _, ok := resultMap[y]; ok {
			result = append(result, y)
			delete(resultMap, y)
		}
	}
	return result
}

func difference(X, Y []string) []string {
	resultMap := make(map[string]bool)
	for _, x := range X {
		resultMap[x] = true
	}
	for _, y := range Y {
		delete(resultMap, y)
	}
	return keys(resultMap)
}

func in(X []string, s string) bool {
	resultMap := make(map[string]bool)
	for _, x := range X {
		resultMap[x] = true
	}
	_, ok := resultMap[s]
	return ok
}

func keys(m map[string]bool) []string {
	result := []string{}
	for k := range m {
		result = append(result, k)
	}
	return result
}
