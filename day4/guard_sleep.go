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

// func (d date) Year() int {
// 	return d.year
// }

// func (d date) Month() int {
// 	return d.month
// }

// func (d date) Day() int {
// 	return d.day
// }

// func (d date) Minute() int {
// 	return d.minute
// }

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
