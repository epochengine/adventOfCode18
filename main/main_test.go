package main

import "testing"

func TestParseStrings(t *testing.T) {
	input := "a\nb\nab\n asdf\n"
	expected := []string{"a", "b", "ab", "asdf"}
	output := ParseStrings(input)

	if len(output) != len(expected) {
		t.Errorf("Output has length %d but expected length of %d", len(output), len(expected))
	}
	for i, s := range expected {
		if output[i] != s {
			t.Errorf("Expected '%s' at location %d but got '%s'", s, i, output[i])
		}
	}
}

func TestParseInts(t *testing.T) {
	input := "1\n2\n-3\n +4\n "
	expected := []int{1, 2, -3, 4}
	output := ParseInts(input)

	if len(output) != len(expected) {
		t.Errorf("Output has length %d but expected length of %d", len(output), len(expected))
	}
	for i, s := range expected {
		if output[i] != s {
			t.Errorf("Expected '%d' at location %d but got '%d'", s, i, output[i])
		}
	}
}
