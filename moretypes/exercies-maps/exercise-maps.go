package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	tokens := strings.Split(s, " ")
	tokenMap := make(map[string]int)

	for _, token := range tokens {
		tokenMap[token]++
	}
	return tokenMap
}

func main() {
	wc.Test(WordCount)
}
