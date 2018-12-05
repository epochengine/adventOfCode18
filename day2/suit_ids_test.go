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

func TestCalculateHammingDifference(t *testing.T) {
	cases := []struct {
		first    string
		second   string
		expected []int
	}{
		{"abcd", "abcd", []int{}},
		{"abcd", "abdc", []int{2, 3}},
		{"abcd", "abce", []int{3}},
		{"abcd", "bbce", []int{0, 3}},
	}

	for _, c := range cases {
		output := CalculateHammingDifference(c.first, c.second)
		if len(output) != len(c.expected) {
			t.Errorf("Output has length %d but expected length of %d", len(output), len(c.expected))
		}

		for i, x := range c.expected {
			if output[i] != x {
				t.Errorf("CalculateHammingDifference for %s and %s expected %d at location %d but got %d", c.first, c.second, x, i, output[i])
			}
		}
	}
}

func TestRemoveCharacterAtIndex(t *testing.T) {
	cases := []struct {
		input    string
		index    int
		expected string
	}{
		{"abcd", 0, "bcd"},
		{"abcd", 1, "acd"},
		{"abcd", 3, "abc"},
		{"abcd", 3, "abc"},
		{"abcdefgh", 4, "abcdfgh"},
	}

	for _, c := range cases {
		output := RemoveCharacterAtIndex(c.input, c.index)
		if output != c.expected {
			t.Errorf("RemoveCharacterAtIndex for string %s and index %d returned %s but expected %s", c.input, c.index, output, c.expected)
		}
	}
}

func TestFindCommonCharactersOfSimilarIds(t *testing.T) {
	cases := []struct {
		input []string
		expected string
	} {
		{[]string{"abcd", "abce", "aacc"}, "abc"},
		{[]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}, "fgij"},
	}

	for _, c := range cases {
		output := FindCommonCharactersOfSimilarIds(c.input)
		if output != c.expected {
			t.Errorf("FindCommonCharactersOfSimilarIds for %v returned %s but expected %s", c.input, output, c.expected)
		}
	}
}
