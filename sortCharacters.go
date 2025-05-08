package main

import (
	"sort"
	"strings"
)

func SortCharacters(word string) string {
	chars := []rune(word)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}

func SortCharacters2(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
