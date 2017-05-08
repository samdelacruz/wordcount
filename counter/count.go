package counter

import "strings"
import "regexp"

// FromString returns the count of each word in a string
func FromString(text *string) map[string]int {
	words := strings.Split(*text, " ")
	return fromWords(&words)
}

func filterWord(word *string) *string {
	w := strings.ToLower(*word)
	reg, _ := regexp.Compile("[^a-z]")
	w = reg.ReplaceAllString(w, "")
	if w == *word {
		return word
	}
	return &w
}

func fromWords(words *[]string) map[string]int {
	counts := make(map[string]int)

	for i := range *words {
		word := filterWord(&(*words)[i])
		if len(*word) > 0 {
			counts[*word]++
		}
	}

	return counts
}
