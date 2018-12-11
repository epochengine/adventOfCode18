package day4

import (
	"regexp"
	"sort"
	"strconv"
)

type scheduleEntry interface {
	date() date
}

type date struct {
	year   int
	month  int
	day    int
	minute int
}

type startShift struct {
	entryDate date
	guardID   int
}

func (ss startShift) date() date {
	return ss.entryDate
}

type fallAsleep struct {
	entryDate date
}

func (fa fallAsleep) date() date {
	return fa.entryDate
}

type wakeUp struct {
	entryDate date
}

func (wu wakeUp) date() date {
	return wu.entryDate
}

var scheduleEntryRegex = regexp.MustCompile(`\[(\d{4})-(\d{2})-(\d{2}) \d{2}:(\d{2})\] (.+)$`)
var startShiftRegex = regexp.MustCompile(`Guard #(\d+) begins shift`)

func parseScheduleEvent(input string) scheduleEntry {
	parsed := scheduleEntryRegex.FindStringSubmatch(input)[1:]
	var dateParts [4]int
	for i := 0; i < 4; i++ {
		intPart, err := strconv.Atoi(parsed[i])
		if err != nil {
			panic(err)
		}
		dateParts[i] = intPart
	}
	entryDate := date{dateParts[0], dateParts[1], dateParts[2], dateParts[3]}

	eventDetails := parsed[4]
	switch eventDetails {
	case "falls asleep":
		return fallAsleep{entryDate}
	case "wakes up":
		return wakeUp{entryDate}
	default:
		idParse := startShiftRegex.FindStringSubmatch(eventDetails)[1]
		guardID, err := strconv.Atoi(idParse)
		if err != nil {
			panic(err)
		}
		return startShift{entryDate, guardID}
	}
}

func parseGuardSchedule(input []string) []scheduleEntry {
	schedule := make([]scheduleEntry, len(input))
	for i, s := range input {
		entry := parseScheduleEvent(s)
		schedule[i] = entry
	}

	sort.Slice(schedule, func(i, j int) bool {
		iDate := schedule[i].date()
		jDate := schedule[j].date()

		if iDate.year != jDate.year {
			return iDate.year < jDate.year
		}
		if iDate.month != jDate.month {
			return iDate.month < jDate.month
		}
		if iDate.day != jDate.day {
			return iDate.day < jDate.day
		}
		if iDate.minute != jDate.minute {
			return iDate.minute < jDate.minute
		}

		return false
	})

	return schedule
}

func calculateSleepTimes(entries []scheduleEntry) map[int]map[int]int {
	sleepTimes := make(map[int]map[int]int)
	var guardID int
	var startSleep int
	for _, entry := range entries {
		switch e := entry.(type) {
		case startShift:
			guardID = e.guardID
		case fallAsleep:
			startSleep = e.entryDate.minute
		case wakeUp:
			sleeps, ok := sleepTimes[guardID]
			if !ok {
				sleeps = make(map[int]int)
			}

			for i := startSleep; i < e.entryDate.minute; i++ {
				count, ok := sleeps[i]
				if !ok {
					sleeps[i] = 1
				} else {
					sleeps[i] = count + 1
				}
			}
			sleepTimes[guardID] = sleeps
		default:
			panic("Unknown entry type!")
		}
	}

	return sleepTimes
}

func calculateTotalSleepingMinutes(sleepTimes map[int]int) int {
	total := 0
	for _, minutes := range sleepTimes {
		total += minutes
	}
	return total
}

// FindMostLikelySleepingGuardMinute parses a list of string guard sleep
// schedule entries and finds the guard that sleeps the most, and the
// minute that guard is most likely (by frequency) to be asleep.
func FindMostLikelySleepingGuardMinute(input []string) (int, int) {
	scheduleEntries := parseGuardSchedule(input)
	sleepTimes := calculateSleepTimes(scheduleEntries)
	maxMinutes := 0
	guardID := -1

	for id, sleeps := range sleepTimes {
		total := calculateTotalSleepingMinutes(sleeps)
		if total > maxMinutes {
			guardID = id
			maxMinutes = total
		}
	}

	sleeps := sleepTimes[guardID]
	var minute int
	var maxFrequency int
	for min, freq := range sleeps {
		if freq > maxFrequency {
			minute = min
			maxFrequency = freq
		}
	}

	return guardID, minute
}

// FindMostFrequentlyAsleepMinuteGuard parses a list of string guard sleep
// schedule entries and finds the guard who is the asleep most often in any
// minute.
func FindMostFrequentlyAsleepMinuteGuard(input []string) (int, int) {
	scheduleEntries := parseGuardSchedule(input)
	sleepTimes := calculateSleepTimes(scheduleEntries)
	maxFrequency := 0
	minute := -1
	guardID := -1

	for id, sleeps := range sleepTimes {
		for min, freq := range sleeps {
			if freq > maxFrequency {
				guardID = id
				minute = min
				maxFrequency = freq
			}
		}
	}

	return guardID, minute
}
