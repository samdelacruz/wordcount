package counter

import "strings"
import "regexp"

// FromString returns the count of each word in a string
func FromString(text *string) map[string]int {
	words := regexp.MustCompile("[^a-z]+").Split(strings.ToLower(*text), -1)
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
