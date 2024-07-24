package goblok

import "strings"

// normalize replaces leet speak characters with their normal counterparts.
func normalize(input string) string {
	var normalized strings.Builder
	for _, char := range input {
		if standardChar, found := LeetMap[char]; found {
			normalized.WriteRune(standardChar)
		} else {
			normalized.WriteRune(char)
		}
	}
	return normalized.String()
}

// findOriginalIndex helps find the index of the original word in the unnormalized input.
func findOriginalIndex(original, normalized string, normalizedIndex int, _ string) int {
	originalIndex := 0
	normLen := 0
	for i := 0; i < len(original); i++ {
		if normLen >= normalizedIndex {
			return originalIndex
		}
		if normalize(string(original[i])) == string(normalized[i]) {
			normLen++
		} else {
			normLen += len(normalize(string(original[i])))
		}
		originalIndex++
	}
	return -1
}
