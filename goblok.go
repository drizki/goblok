package goblok

import (
	"errors"
	"strings"
)

type Goblok struct {
	ProfaneWords []string
}

type Analysis struct {
	Input     string
	Profane   []string
	Filtered  string
	Locations []Location
}

type Location struct {
	Word  string
	Index int
}

// New creates a new instance of Goblok. Users can provide their own dictionary of profane words.
func New(dictionary ...[]string) *Goblok {
	profaneWords := ProfaneWords
	if len(dictionary) > 0 {
		profaneWords = dictionary[0]
	}
	return &Goblok{ProfaneWords: profaneWords}
}

// Add adds a new profane word to the dictionary.
func (goblok *Goblok) Add(word string) string { 
	if word == "" {
		return ""
	} 
	for _, w := range goblok.ProfaneWords {
		if w == word {
			return ""
		}
	}

	goblok.ProfaneWords = append(goblok.ProfaneWords, word)
	return word
}

// Remove removes a profane word from the dictionary.
func (goblok *Goblok) Remove(word string) string { 
	if word == "" {
		return ""
	} 
	for i, w := range goblok.ProfaneWords {
		if w == word {
			goblok.ProfaneWords = append(goblok.ProfaneWords[:i], goblok.ProfaneWords[i+1:]...)
			return word
		}
	}

	return ""
}

// Filter applies a profanity filter to the input string, replacing profane words with a mask character.
func (goblok *Goblok) Filter(input string, mask ...rune) string {
	if len(mask) == 0 {
		mask = []rune{'*'}
	}

	normalizedInput := normalize(input)
	inputLower := strings.ToLower(input)
	filterMap := make(map[int]int)

	for _, word := range goblok.ProfaneWords {
		normalizedWord := normalize(word)
		wordLen := len(word)

		for i := 0; i < len(inputLower); i++ {
			if strings.HasPrefix(inputLower[i:], word) {
				filterMap[i] = wordLen
			}
		}
		for i := 0; i < len(normalizedInput); i++ {
			if strings.HasPrefix(normalizedInput[i:], normalizedWord) {
				filterMap[i] = wordLen
			}
		}
	}

	var result strings.Builder
	for i := 0; i < len(input); {
		if length, found := filterMap[i]; found {
			result.WriteString(strings.Repeat(string(mask), length))
			i += length
		} else {
			result.WriteByte(input[i])
			i++
		}
	}

	return result.String()
}

// Analyze analyzes the input string and returns an Analysis struct containing information about the input.
func (goblok *Goblok) Analyze(input string) (Analysis, error) {
	if input == "" {
		return Analysis{}, errors.New("empty string passed")
	}

	normalizedInput := normalize(strings.ToLower(input))
	profane := []string{}
	locations := []Location{}

	for _, word := range goblok.ProfaneWords {
		normalizedWord := normalize(strings.ToLower(word))
		if strings.Contains(normalizedInput, normalizedWord) {
			profane = append(profane, word) 
			for start := 0; start < len(normalizedInput); {
				index := strings.Index(normalizedInput[start:], normalizedWord)
				if index == -1 {
					break
				} 
				actualIndex := findOriginalIndex(input, normalizedInput, start+index, normalizedWord)
				if actualIndex != -1 {
					locations = append(locations, Location{Word: word, Index: actualIndex})
				}
				start += index + len(normalizedWord)
			}
		}
	}

	filtered := goblok.Filter(input)

	return Analysis{
		Input:     input,
		Filtered:  filtered,
		Profane:   profane,
		Locations: locations,
	}, nil
}

// Contains checks if the input string contains any profane words.
func (goblok *Goblok) Contains(input string) bool {
	analysis, _ := goblok.Analyze(input)
	return len(analysis.Profane) > 0
}
