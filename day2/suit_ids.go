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

// CalculateHammingDifference calculates the indices of differing characters
// between the two given strings, i.e. abc and acb have a Hamming difference
// of [1, 2].
// Only well defined for strings of the same length.
// The Hamming *distance* is the length of the returned slice.
func CalculateHammingDifference(first string, second string) []int {
	difference := []int{}
	for i, c := range first {
		if c != rune(second[i]) {
			difference = append(difference, i)
		}
	}

	return difference
}

// RemoveCharacterAtIndex removes a single character at the given index from
// a string and returns the remaining string.
func RemoveCharacterAtIndex(input string, index int) string {
	return input[0:index] + input[index+1:len(input)]
}

// FindCommonCharactersOfSimilarIds runs through all combinations of the given
// input strings, finds a pair that only have one character different and
// returns the common characters between those two strings.
func FindCommonCharactersOfSimilarIds(input []string) string {
	for i, s1 := range input {
		for _, s2 := range input[i+1:] {
			difference := CalculateHammingDifference(s1, s2)
			if len(difference) == 1 {
				return RemoveCharacterAtIndex(s1, difference[0])
			}
		}
	}

	panic("No matching solution found!")
}
