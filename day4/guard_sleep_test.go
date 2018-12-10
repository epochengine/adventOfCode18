package day4

import "testing"

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
		{"[1519-12-02 00:21] wakes up", wakeUp{date{1519, 12, 2, 21}}},
	}

	for _, c := range cases {
		event := parseScheduleEvent(c.input)
		if event != c.expected {
			t.Errorf("ParseSleepEvent returned %v for %s, but expected %v", event, c.input, c.expected)
		}
	}
}
