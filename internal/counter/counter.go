package counter

import (
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

func CountWords(input string) []WordCount {
	// Split the input string into words
	words := strings.Fields(input)

	// Create a map to store word counts
	wordCounts := make(map[string]int)

	// Count the occurrences of each word
	for _, word := range words {
		wordCounts[word]++
	}

	// Convert the map to a slice of WordCount structs
	var wordCountList []WordCount
	for word, count := range wordCounts {
		wordCountList = append(wordCountList, WordCount{word, count})
	}

	return wordCountList
}

func MostFrequentWords(input string, num int) []WordCount {
	wordCountList := CountWords(input)

	// Sort the WordCount slice by count in descending order
	sortedWordCount := sortByCountDescending(wordCountList)

	// Return the top 'num' most frequent words
	if num > len(sortedWordCount) {
		return sortedWordCount
	}

	return sortedWordCount[:num]
}

func sortByCountDescending(wordCountList []WordCount) []WordCount {
	// Custom sorting function to sort WordCount by count in descending order
	sort.Slice(wordCountList, func(i, j int) bool {
		return wordCountList[i].Count > wordCountList[j].Count
	})
	return wordCountList
}
