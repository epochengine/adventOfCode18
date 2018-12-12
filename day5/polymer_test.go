package day5

import "testing"

func TestRemoveReactingPairs(t *testing.T) {
	cases := []struct {
		polymer string
		reacted string
	}{
		{"aA", ""},
		{"abBA", ""},
		{"abAB", "abAB"},
		{"aabAAB", "aabAAB"},
		{"dabAcCaCBAcCcaDA", "dabCBAcaDA"},
	}

	for _, c := range cases {
		reacted := RemoveReactingPairs(c.polymer)
		if reacted != c.reacted {
			t.Errorf("RemoveReactingPairs calculated %s for input %s but expected %s", reacted, c.polymer, c.reacted)
		}
	}
}

func TestRemoveUnitFromPolymer(t *testing.T) {
	cases := []struct {
		polymer string
		unit    string
		removed string
	}{
		{"dabAcCaCBAcCcaDA", "a", "dbcCCBcCcD"},
		{"dabAcCaCBAcCcaDA", "b", "daAcCaCAcCcaDA"},
		{"dabAcCaCBAcCcaDA", "c", "dabAaBAaDA"},
		{"dabAcCaCBAcCcaDA", "d", "abAcCaCBAcCcaA"},
	}

	for _, c := range cases {
		removed := removeUnitFromPolymer(c.polymer, c.unit)
		if removed != c.removed {
			t.Errorf("RemoveUnitFromPolymer found %s for polymer %s and unit %s but expected %s", removed, c.polymer, c.unit, c.removed)
		}
	}
}

func TestFindShortestPolymerByRemovingUnit(t *testing.T) {
	polymer := "dabAcCaCBAcCcaDA"
	expected := "daDA"
	shortestReacted := FindShortestPolymerByRemovingUnit(polymer)
	if shortestReacted != expected {
		t.Errorf("FindShortestPolymerByRemovingUnit found %s for input %s but expected %s", shortestReacted, polymer, expected)
	}
}
