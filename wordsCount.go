package main

import (
	"fmt"
	"strings"
)

func countWords() {
	text := "Cosmo,is,an,incredible,technical,companion,with,very,strong,skills,in,Algorithms,and,Data,Structures,and,a,great,teacher,for,technical,interviews."

	words := strings.Split(text, ",")

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println(words)
}
