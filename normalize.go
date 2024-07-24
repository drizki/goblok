package goblok

import "strings"

var leetMap = map[rune]rune{
	'0': 'o', '1': 'i', '2': 'z', '3': 'e', '4': 'a', '5': 's', '6': 'g', '7': 't', '8': 'b', '9': 'g',
	'!': 'i', '@': 'a', '(': 'c', ')': 'c', '[': 'c', ']': 'c', '{': 'c', '}': 'c', '^': 'i', '*': 's',
	'~': 'n', 'û': 'u', 'µ': 'u', '€': 'e', '£': 'l', '¥': 'y', '¢': 'c', '₩': 'w', '₱': 'p', '₹': 'r',
	'₽': 'r', '₿': 'b', '฿': 'b', '₺': 'l', '₸': 't', '₮': 't', '₦': 'n', '₴': 'h', '₲': 'g', '₭': 'k',
	'₯': 'd', '₠': 'e', '₡': 'c', '₢': 'c', '₣': 'f', '₤': 'l', '₥': 'm', '₧': 'p', '₨': 'r', '₪': 'i',
	'#': 'h', '$': 's', '%': 'o', '&': 'n', '-': 't', '_': 't', '+': 't', '=': 't', ';': 'i', ':': 'i',
	',': 'i', '<': 'c', '>': 'c',
}


// normalize replaces leet speak characters with their normal counterparts.
func normalize(input string) string {
	var normalized strings.Builder
	for _, char := range input {
		if standardChar, found := leetMap[char]; found {
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
