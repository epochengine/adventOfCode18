package day6

import (
	"testing"
)

func TestParseCoordinates(t *testing.T) {
	cases := []struct {
		input string
		coord coordinate
	}{
		{"1, 1", coordinate{1, 1}},
		{"10, 1", coordinate{10, 1}},
		{"2, 4", coordinate{2, 4}},
		{"11, 999", coordinate{11, 999}},
	}

	for _, c := range cases {
		coord := parseCoordinates(c.input)
		if coord != c.coord {
			t.Errorf("ParseCoordinates returned %v for input %s but expected %v", coord, c.input, c.coord)
		}
	}
}

func TestCalculateManhattenDistance(t *testing.T) {
	cases := []struct {
		first    coordinate
		second   coordinate
		distance int
	}{
		{coordinate{1, 1}, coordinate{1, 2}, 1},
		{coordinate{1, 1}, coordinate{2, 1}, 1},
		{coordinate{1, 1}, coordinate{2, 2}, 2},
		{coordinate{1, 1}, coordinate{4, 4}, 6},
		{coordinate{2, 3}, coordinate{3, 1}, 3},
		{coordinate{3, 3}, coordinate{2, 4}, 2},
		{coordinate{3, 4}, coordinate{1, 2}, 4},
	}

	for _, c := range cases {
		distance := calculateManhattenDistance(c.first, c.second)
		if distance != c.distance {
			t.Errorf("CalculateManhattenDistance returned %d for coordinates %v and %v, but expected %d", distance, c.first, c.second, c.distance)
		}
	}
}

func TestCoordinateMapSort(t *testing.T) {
	coordinates := []coordinate{
		coordinate{1, 1},
		coordinate{1, 6},
		coordinate{1, 3},
		coordinate{8, 3},
		coordinate{3, 4},
		coordinate{5, 5},
		coordinate{8, 9},
	}

	expected := []coordinate{
		coordinate{1, 1},
		coordinate{1, 3},
		coordinate{1, 6},
		coordinate{3, 4},
		coordinate{5, 5},
		coordinate{8, 3},
		coordinate{8, 9},
	}

	coordMap := coordinateMap{coordinates}
	coordMap.sort()

	for i, coord := range expected {
		if coordMap.coordinates[i] != coord {
			t.Errorf("Sorted coordinate map at index i had %v but expected %v", coordMap.coordinates[i], coord)
		}
	}
}

func TestCoordinateMapFindEdges(t *testing.T) {
	coordinates := []coordinate{
		coordinate{3, 4},
		coordinate{1, 1},
		coordinate{1, 6},
		coordinate{8, 3},
		coordinate{5, 5},
		coordinate{8, 9},
	}

	expectedMinX := 1
	expectedMinY := 1
	expectedMaxX := 8
	expectedMaxY := 9

	coordMap := coordinateMap{coordinates}

	minX := coordMap.minX()
	if minX != expectedMinX {
		t.Errorf("MinX returned %d but expected %d", minX, expectedMinX)
	}
	minY := coordMap.minY()
	if minY != expectedMinY {
		t.Errorf("MinY returned %d but expected %d", minY, expectedMinY)
	}
	maxX := coordMap.maxX()
	if maxX != expectedMaxX {
		t.Errorf("MaxX returned %d but expected %d", maxX, expectedMaxX)
	}
	maxY := coordMap.maxY()
	if maxY != expectedMaxY {
		t.Errorf("MaxY returned %d but expected %d", maxY, expectedMaxY)
	}
}
