package day4

import (
	"reflect"
	"testing"
)

func TestParseGuardSchedule(t *testing.T) {
	input := []string{
		"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-02 00:40] falls asleep",
		"[1518-11-01 00:55] wakes up",
		"[1518-11-01 00:05] falls asleep",
		"[1518-11-02 00:50] wakes up",
		"[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:30] falls asleep",
		"[1518-11-01 23:58] Guard #99 begins shift",
	}
	expected := []scheduleEntry{
		startShift{date{1518, 11, 1, 0}, 10},
		fallAsleep{date{1518, 11, 1, 5}},
		wakeUp{date{1518, 11, 1, 25}},
		fallAsleep{date{1518, 11, 1, 30}},
		wakeUp{date{1518, 11, 1, 55}},
		startShift{date{1518, 11, 1, 58}, 99},
		fallAsleep{date{1518, 11, 2, 40}},
		wakeUp{date{1518, 11, 2, 50}},
	}
	schedule := parseGuardSchedule(input)

	for i, entry := range expected {
		if schedule[i] != entry {
			t.Errorf("Expected entry %v at position %d but got %v", entry, i, schedule[i])
		}
	}
}

func TestParseScheduleEvent(t *testing.T) {
	cases := []struct {
		input    string
		expected scheduleEntry
	}{
		{"[1518-11-01 00:00] Guard #10 begins shift", startShift{date{1518, 11, 1, 0}, 10}},
		{"[1518-11-03 00:30] falls asleep", fallAsleep{date{1518, 11, 3, 30}}},
		{"[1519-12-31 00:21] wakes up", wakeUp{date{1519, 12, 31, 21}}},
	}

	for _, c := range cases {
		event := parseScheduleEvent(c.input)
		if event != c.expected {
			t.Errorf("ParseSleepEvent returned %v for %s, but expected %v", event, c.input, c.expected)
		}
	}
}

func TestCalculateSleepTimes(t *testing.T) {
	entries := []scheduleEntry{
		startShift{date{1518, 11, 1, 0}, 1},
		fallAsleep{date{1518, 11, 1, 10}},
		wakeUp{date{1518, 11, 1, 12}},
		fallAsleep{date{1518, 11, 1, 27}},
		wakeUp{date{1518, 11, 1, 30}},
		startShift{date{1518, 11, 2, 0}, 2},
		fallAsleep{date{1518, 11, 2, 1}},
		wakeUp{date{1518, 11, 2, 5}},
		startShift{date{1518, 11, 3, 0}, 1},
		fallAsleep{date{1518, 11, 2, 11}},
		wakeUp{date{1518, 11, 2, 12}},
	}

	expected := map[int]map[int]int{
		1: map[int]int{10: 1, 11: 2, 27: 1, 28: 1, 29: 1},
		2: map[int]int{1: 1, 2: 1, 3: 1, 4: 1},
	}

	sleepTimes := calculateSleepTimes(entries)
	if !reflect.DeepEqual(expected, sleepTimes) {
		t.Errorf("CalculateSleepTimes returned %v but expected %v", sleepTimes, expected)
	}
}

func TestCalculateTotalSleepingMinutes(t *testing.T) {
	cases := []struct {
		sleepTimes map[int]int
		total int
	} {
		{map[int]int{10: 1, 11: 2, 27: 1, 28: 1, 29: 1}, 6},
		{map[int]int{1: 1, 2: 1, 3: 2, 4: 3}, 7},
	}
	
	for _, c := range cases {
		foundTotal := calculateTotalSleepingMinutes(c.sleepTimes)
		if foundTotal != c.total {
			t.Errorf("CalculateTotalSleepingMinutes found total %d for input %v but expected %d", foundTotal, c.sleepTimes, c.total)
		}
	}
}

func TestFindMostLikelySleepingGuardMinute(t *testing.T) {
	input := []string{
		"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-01 00:05] falls asleep",
		"[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:30] falls asleep",
		"[1518-11-01 00:55] wakes up",
		"[1518-11-01 23:58] Guard #99 begins shift",
		"[1518-11-02 00:40] falls asleep",
		"[1518-11-02 00:50] wakes up",
		"[1518-11-03 00:05] Guard #10 begins shift",
		"[1518-11-03 00:24] falls asleep",
		"[1518-11-03 00:29] wakes up",
		"[1518-11-04 00:02] Guard #99 begins shift",
		"[1518-11-04 00:36] falls asleep",
		"[1518-11-04 00:46] wakes up",
		"[1518-11-05 00:03] Guard #99 begins shift",
		"[1518-11-05 00:45] falls asleep",
		"[1518-11-05 00:55] wakes up",
	}
	expectedID, expectedMinute := 10, 24

	id, minute := FindMostLikelySleepingGuardMinute(input)
	if id != expectedID {
		t.Errorf("FindMostLikelySleepingGuardMinute found ID %d but expected %d", id, expectedID)
	}
	if minute != expectedMinute {
		t.Errorf("FindMostLikelySleepingGuardMinute found minute %d but expected %d", minute, expectedMinute)
	}
}
