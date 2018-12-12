package main

import (
	"epochengine/adventofcode/day1"
	"epochengine/adventofcode/day2"
	"epochengine/adventofcode/day3"
	"epochengine/adventofcode/day4"
	"epochengine/adventofcode/day5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	calculateDay5()
}

func calculateDay1() {
	rawInput, err := Asset("resources/frequency_input.txt")
	if err != nil {
		panic(err)
	}
	input := ParseInts(string(rawInput))

	fmt.Printf("Day 1 part 1: %d\n", day1.CalculateFrequency(input))
	fmt.Printf("Day 1 part 2: %d\n", day1.FindFirstRepeatFrequency(input))
}

func calculateDay2() {
	rawInput, err := Asset("resources/suit_ids_input.txt")
	if err != nil {
		panic(err)
	}
	input := ParseStrings(string(rawInput))

	fmt.Printf("Day 2 part 1: %d\n", day2.CalculateChecksum(input))
	fmt.Printf("Day 2 part 2: %s\n", day2.FindCommonCharactersOfSimilarIds(input))
}

func calculateDay3() {
	rawInput, err := Asset("resources/fabric_input.txt")
	if err != nil {
		panic(err)
	}
	input := ParseStrings(string(rawInput))

	fmt.Printf("Day 3 part 1: %d\n", day3.FindOverlappingSquares(input))
	fmt.Printf("Day 3 part 2: %d\n", day3.FindUniqueClaim(input))
}

func calculateDay4() {
	rawInput, err := Asset("resources/guard_sleep_input.txt")
	if err != nil {
		panic(err)
	}
	input := ParseStrings(string(rawInput))

	id, minute := day4.FindMostLikelySleepingGuardMinute(input)
	fmt.Printf("Day 4 part 1: %d * %d = %d\n", id, minute, id*minute)
	id, minute = day4.FindMostFrequentlyAsleepMinuteGuard(input)
	fmt.Printf("Day 4 part 2: %d * %d = %d\n", id, minute, id*minute)
}

func calculateDay5() {
	rawInput, err := Asset("resources/polymer_input.txt")
	if err != nil {
		panic(err)
	}
	input := string(rawInput)
	input = strings.TrimSuffix(input, "\n")

	reacted := day5.RemoveReactingPairs(input)
	fmt.Printf("Day 5 part 1: %d\n", len(reacted))
	shortestReacted := day5.FindShortestPolymerByRemovingUnit(input)
	fmt.Printf("Day 5 part 2: %d\n", len(shortestReacted))
}

// ParseStrings takes a newline separated single string and creates a slice
// of the individual lines as elements, trimmed of leading/trailing whitespace.
func ParseStrings(input string) []string {
	input = strings.TrimSpace(input)
	input = strings.TrimSuffix(input, "\n")
	split := strings.Split(input, "\n")
	for i, s := range split {
		split[i] = strings.TrimSpace(s)
	}

	return split
}

// ParseInts takes a newline separated single string and creates a slice
// of the individual lines as ints.
func ParseInts(input string) []int {
	splitStrings := ParseStrings(input)
	split := make([]int, len(splitStrings))
	for i, x := range splitStrings {
		converted, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		split[i] = converted
	}

	return split
}
