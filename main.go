package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)
	for _, v := range words {
		wordMap[v]++
	}
	return wordMap
}

func main() {
	fmt.Println(wordCount("Hello, this is my word count app"))
	fmt.Println(wordCount("Hi there, I am just creating this so I can repeat the I word"))
}
