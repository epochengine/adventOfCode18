package main

import (
	"epochengine/adventofcode/day1"
	"fmt"
)

func main() {
	frequencyInput, err := Asset("resources/frequency_input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(frequencyInput)
	fmt.Printf("Day 1 part 1: %d\n", day1.CalculateFrequency(inputStr, "\n"))
	fmt.Printf("Day 1 part 2: %d\n", day1.FindFirstRepeatFrequency(inputStr, "\n"))
}
