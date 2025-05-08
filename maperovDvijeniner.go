package main

import "strings"

func FindMajorityElement(arr []int) int {
	countMap := make(map[int]int)     // Initialize a map to count occurrences of each number
	majorityThreshold := len(arr) / 2 // Define the majority threshold // kara lini cankacac qanakic heto karevorutyun xosi len(arr) / 4

	for _, num := range arr {
		countMap[num]++                        // Increment the count for the current number
		if countMap[num] > majorityThreshold { // Check if current number exceeds majority threshold
			return num // Return the number if it's found to be the majority
		}
	}
	return -1 // Return -1 if no majority element is found
}

func CreateKeywordIndex(docs []string) map[string][]int {
	index := make(map[string][]int) // Initialize a map to store words and their document indices

	for i, doc := range docs {
		words := strings.Fields(doc) // Split the document into individual words

		for _, word := range words {
			index[word] = append(index[word], i) // Append the document index to the word's list
		}
	}
	return index // Return the constructed index map
}
