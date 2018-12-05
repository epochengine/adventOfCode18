package day2

// CalculateCharacterCounts reads a string and populates a map with the counts
// of each character.
func CalculateCharacterCounts(input string) map[rune]int {
	counts := make(map[rune]int)
	for _, c := range input {
		count, ok := counts[c]
		if ok {
			counts[c] = count + 1
		} else {
			counts[c] = 1
		}
	}
	return counts
}

// HasCharacterCount validates if the given map of character counts has any
// character with a count of 'count'.
func HasCharacterCount(characterCounts map[rune]int, count int) bool {
	for _, c := range characterCounts {
		if c == count {
			return true
		}
	}

	return false
}

// CalculateChecksum takes a slice of strings and calculates a rudimentary
// checksum based on the character counts of each string. Only strings that
// have exactly two and/or exactly three of a character are used.
func CalculateChecksum(suitIds []string) int {
	twos, threes := 0, 0
	for _, s := range suitIds {
		characterCounts := CalculateCharacterCounts(s)
		if HasCharacterCount(characterCounts, 2) {
			twos++
		}
		if HasCharacterCount(characterCounts, 3) {
			threes++
		}
	}

	return twos * threes
}
