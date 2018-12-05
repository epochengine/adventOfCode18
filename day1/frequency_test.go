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

func TestCalculateFrequency(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, -2, 3, 1}, 3},
		{[]int{1, 1, 1}, 3},
		{[]int{1, 1, -2}, 0},
		{[]int{-1, -2, -3}, -6},
	}

	for _, c := range cases {
		outputFrequency := CalculateFrequency(c.input)
		if outputFrequency != c.expected {
			t.Errorf("CalculateFrequency returned incorrect result. Got %d but expected %d", outputFrequency, c.expected)
		}
	}
}

func TestFindFirstRepeatFrequency(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{[]int{+1, -2, +3, +1}, 2},
		{[]int{1, -1}, 0},
		{[]int{7, -6}, 7},
		{[]int{3, 3, 4, -2, -4}, 10},
		{[]int{-6, 3, 8, 5, -6}, 5},
		{[]int{7, 7, -2, -7, -4}, 14},
	}

	for _, c := range cases {
		firstRepeat := FindFirstRepeatFrequency(c.input)
		if firstRepeat != c.expected {
			t.Errorf("FindFirstRepeatFrequency returned incorrect result. Got %d but expected %d", firstRepeat, c.expected)
		}
	}
}
