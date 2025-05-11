package maps

func FindAnagramsEfficient(array1, array2 []string) []string {
	sortedWordsInArray2 := make(map[string]bool)
	for _, word := range array2 {
		sortedWordsInArray2[SortCharacters(word)] = true
	}

	result := []string{}

	anagramMatched := make(map[string]bool)

	for _, word := range array1 {
		sortedWord := SortCharacters(word)

		if sortedWordsInArray2[sortedWord] && !anagramMatched[word] {
			result = append(result, word)
			anagramMatched[word] = true
		}
	}

	return result
}

func FindLastUniqueWordEfficient(words []string) string {
	wordsMap := make(map[string]int)

	for _, word := range words {
		wordsMap[word]++
	}

	for i := len(words) - 1; i >= 0; i-- {
		if wordsMap[words[i]] == 1 {
			return words[i]
		}
	}

	return ""
}
