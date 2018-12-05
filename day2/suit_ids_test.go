package day2

import "testing"

func TestCalculateCharacterCounts(t *testing.T) {
	cases := []struct {
		input    string
		expected map[rune]int
	}{
		{"abcdef", map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1}},
		{"bababc", map[rune]int{'a': 2, 'b': 3, 'c': 1}},
		{"abbcde", map[rune]int{'a': 1, 'b': 2, 'c': 1, 'd': 1, 'e': 1}},
		{"abcccd", map[rune]int{'a': 1, 'b': 1, 'c': 3, 'd': 1}},
		{"aabcdd", map[rune]int{'a': 2, 'b': 1, 'c': 1, 'd': 2}},
		{"abcdee", map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 2}},
		{"ababab", map[rune]int{'a': 3, 'b': 3}},
	}

	for _, c := range cases {
		output := CalculateCharacterCounts(c.input)
		for character, count := range c.expected {
			outputCount, ok := output[character]
			if !ok {
				t.Errorf("CalculateCharacterCounts for %s did not contain expected letter %v", c.input, character)
			}
			if outputCount != count {
				t.Errorf("CalculateCharacterCounts for %s had count %d for letter %v but expected %d", c.input, outputCount, character, count)
			}
		}
	}
}

func TestHasCharacterCount(t *testing.T) {
	cases := []struct {
		input    map[rune]int
		count    int
		expected bool
	}{
		{map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1}, 1, true},
		{map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1}, 2, false},
		{map[rune]int{'a': 2, 'b': 3, 'c': 1}, 2, true},
		{map[rune]int{'a': 2, 'b': 3, 'c': 1}, 3, true},
	}

	for _, c := range cases {
		output := HasCharacterCount(c.input, c.count)
		if output != c.expected {
			t.Errorf("HasCharacterCount for %v and count %d should be %v but was %v", c.input, c.count, c.expected, output)
		}
	}
}

func TestCalculateChecksum(t *testing.T) {
	cases := []struct {
		input    []string
		expected int
	}{
		{[]string{"a", "aa"}, 0},
		{[]string{"aa", "bb"}, 0},
		{[]string{"aaa", "aa"}, 1},
		{[]string{"abab", "aabbb"}, 2},
		{[]string{"aa", "aaabbb"}, 1},
		{[]string{"aabb", "aaabbb"}, 1},
		{[]string{"aabb", "aaabbb", "ccc"}, 2},
		{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, 12},
	}

	for _, c := range cases {
		output := CalculateChecksum(c.input)
		if output != c.expected {
			t.Errorf("CalculateChecksum for %v generated checksum of %d but expected %d", c.input, output, c.expected)
		}
	}
}
