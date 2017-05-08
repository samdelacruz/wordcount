package wordcount

import "strings"

// FromString counts the number of occurences of each word in a string
func FromString(text *string) map[string]int {
	words := strings.Split(strings.ToLower(*text), " ")
	return FromWords(&words)
}

// FromWords counts the number of occurences of each word in a list of strings
func FromWords(words *[]string) map[string]int {
	counts := make(map[string]int)

	for _, word := range *words {
		counts[word]++
	}

	return counts
}
