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
	inputs := make([]int, len(stringInputs))
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

// CalculateFrequency will read the input and calculate the resulting frequency
// after applying all the adjustments, starting from
func CalculateFrequency(input string) int {
	frequencies := ParseInput(input, "\n")
	frequency := 0
	for _, x := range frequencies {
		frequency = AdjustFrequency(frequency, x)
	}

	return frequency
}
