package day1

import "testing"

func TestAdjustFrequency(t *testing.T) {
	cases := []struct {
		start      int
		adjustment int
		expected   int
	}{
		{1, 1, 2},
		{0, 1, 1},
		{-1, -1, -2},
		{0, -1, -1},
		{10000, -1, 9999},
	}

	for _, c := range cases {
		output := AdjustFrequency(c.start, c.adjustment)
		if output != c.expected {
			t.Errorf("AdjustFrequency returned %d, expected %d", output, c.expected)
		}
	}
}

func TestParseInput(t *testing.T) {
	cases := []struct {
		input     string
		expected  []int
		splitChar string
	}{
		{"+1, -1, 200, -11, +12", []int{1, -1, 200, -11, 12}, ","},
		{"100\n-1", []int{100, -1}, "\n"},
	}

	for _, c := range cases {
		output := ParseInput(c.input, c.splitChar)

		for i, x := range output {
			if x != c.expected[i] {
				t.Errorf("ParseInput did not return the correct output, got %d and expected %d", x, c.expected[i])
			}
		}
	}
}
