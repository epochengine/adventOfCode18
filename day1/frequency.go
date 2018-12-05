package day1

// AdjustFrequency raises/lowers the input frequency by the adjustment provided.
func AdjustFrequency(frequency int, adjustment int) int {
	return frequency + adjustment
}

// CalculateFrequency reads the input and calculates the resulting frequency
// after applying all the adjustments, starting from 0.
func CalculateFrequency(frequencies []int) int {
	frequency := 0
	for _, x := range frequencies {
		frequency = AdjustFrequency(frequency, x)
	}

	return frequency
}

// FindFirstRepeatFrequency reads the input and continuously adjusts the
// frequency until the first repeated adjusted frequency is found.
func FindFirstRepeatFrequency(frequencies []int) int {
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
