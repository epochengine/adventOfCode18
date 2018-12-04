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

func TestCalculateFrequency(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"+1, -2, +3, +1", 3},
		{"1, 1, 1", 3},
		{"1, 1, -2", 0},
		{"-1, -2, -3", -6},
	}

	for _, c := range cases {
		outputFrequency := CalculateFrequency(c.input, ",")
		if outputFrequency != c.expected {
			t.Errorf("CalculateFrequency returned incorrect result. Got %d but expected %d", outputFrequency, c.expected)
		}
	}
}

func TestFindFirstRepeatFrequency(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"+1, -2, +3, +1", 2},
		{"1, -1", 0},
		{"7, -6", 7},
		{"3, 3, 4, -2, -4", 10},
		{"-6, 3, 8, 5, -6", 5},
		{"7, 7, -2, -7, -4", 14},
	}

	for _, c := range cases {
		firstRepeat := FindFirstRepeatFrequency(c.input, ",")
		if firstRepeat != c.expected {
			t.Errorf("FindFirstRepeatFrequency returned incorrect result. Got %d but expected %d", firstRepeat, c.expected)
		}
	}
}
