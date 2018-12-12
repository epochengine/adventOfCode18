package day5

import (
	"strings"
)

// RemoveReactingPairs reads in a polymer (string) and iterates through the
// characters, removing any pairs of opposite-case characters (e.g. aA)
// until none remain.
func RemoveReactingPairs(polymer string) string {
	reacted := polymer
	diff := 'a' - 'A'

	for {
		finished := true
		var lastChar rune

		for pos, char := range reacted {
			if pos == 0 {
				lastChar = char
				continue
			}

			if char-lastChar == diff || lastChar-char == diff {
				finished = false
				reacted = reacted[:pos-1] + reacted[pos+1:]
				break
			}
			lastChar = char
		}

		if finished {
			break
		}
	}

	return reacted
}

func removeUnitFromPolymer(polymer string, unit string) string {
	upper := strings.ToUpper(unit)
	lower := strings.ToLower(unit)
	polymer = strings.Replace(polymer, upper, "", -1)
	polymer = strings.Replace(polymer, lower, "", -1)
	return polymer
}

// FindShortestPolymerByRemovingUnit reads in a polymer (string) and iterates
// removing each possible 'unit' (character, both upper- and lower-case) and
// removing reacting pairs, finding the shortest resulting polymer.
func FindShortestPolymerByRemovingUnit(polymer string) string {
	units := make(map[string]struct{})
	for _, char := range polymer {
		lower := strings.ToLower(string(char))
		units[lower] = struct{}{}
	}

	shortest := ""
	shortestLength := len(polymer)

	for unit := range units {
		removedPolymer := removeUnitFromPolymer(polymer, unit)
		reacted := RemoveReactingPairs(removedPolymer)
		length := len(reacted)
		if length < shortestLength {
			shortest = reacted
			shortestLength = length
		}
	}

	return shortest
}
