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
	fmt.Printf("%d\n", day1.CalculateFrequency(string(frequencyInput)))
}
