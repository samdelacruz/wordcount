package counter

import "strings"

// FromString returns the count of each word in a string
func FromString(text *string) map[string]int {
	words := strings.Split(strings.ToLower(*text), " ")
	return fromWords(&words)
}

func fromWords(words *[]string) map[string]int {
	counts := make(map[string]int)

	for i := range *words {
		word := &(*words)[i]
		if len(*word) > 0 {
			counts[*word]++
		}
	}

	return counts
}
