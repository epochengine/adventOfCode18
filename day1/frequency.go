package day1

import (
	"strconv"
	"strings"
)

// AdjustFrequency raises/lowers the input frequency by the adjustment provided.
func AdjustFrequency(frequency int, adjustment int) int {
	return frequency + adjustment
}

// ParseInput reads a string of ints and creates a slice of their values.
func ParseInput(input string, splitChar string) []int {
	stringInputs := strings.Split(input, splitChar)
	inputs := make([]int, len(stringInputs)-1)
	for i, x := range stringInputs {
		x = strings.TrimSpace(x)
		if len(x) == 0 {
			continue
		}

		input, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		inputs[i] = input
	}

	return inputs
}

// CalculateFrequency reads the input and calculates the resulting frequency
// after applying all the adjustments, starting from 0.
func CalculateFrequency(input string, splitChar string) int {
	frequencies := ParseInput(input, splitChar)
	frequency := 0
	for _, x := range frequencies {
		frequency = AdjustFrequency(frequency, x)
	}

	return frequency
}

// FindFirstRepeatFrequency reads the input and continuously adjusts the
// frequency until the first repeated adjusted frequency is found.
func FindFirstRepeatFrequency(input string, splitChar string) int {
	frequencies := ParseInput(input, splitChar)
	frequency := 0
	seenFrequencies := make(map[int]struct{})
	seenFrequencies[0] = struct{}{}

	for {
		for _, x := range frequencies {
			frequency = AdjustFrequency(frequency, x)
			_, ok := seenFrequencies[frequency]
			if ok {
				return frequency
			}
			seenFrequencies[frequency] = struct{}{}
		}
	}
}
