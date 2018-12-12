package day6

import (
	"sort"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func parseCoordinates(input string) coordinate {
	coords := strings.Split(input, ",")
	parts := make([]int, 2)
	for i, coord := range coords {
		part, err := strconv.Atoi(strings.TrimSpace(coord))
		if err != nil {
			panic(err)
		}
		parts[i] = part
	}

	return coordinate{parts[0], parts[1]}
}

func calculateManhattenDistance(first coordinate, second coordinate) int {
	return abs(first.x-second.x) + abs(first.y-second.y)
}

// Two's complement abs for ints to avoid conversion.
// Courtesy of http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
func abs(n int) int {
	y := n >> 31
	return (n ^ y) - y
}

type coordinateMap struct {
	coordinates []coordinate
}

func (cm *coordinateMap) sort() {
	sort.Slice(cm.coordinates, func(i, j int) bool {
		first, second := cm.coordinates[i], cm.coordinates[j]
		if first.x < second.x {
			return true
		} else if first.x == second.x && first.y < second.y {
			return true
		}
		return false
	})
}

func (cm coordinateMap) minX() int {
	if cm.coordinates == nil || len(cm.coordinates) == 0 {
		panic("Cannot find min x for empty/nil map.")
	}
	minX := cm.coordinates[0].x
	for i := 1; i < len(cm.coordinates); i++ {
		x := cm.coordinates[i].x
		if x < minX {
			minX = x
		}
	}

	return minX
}

func (cm coordinateMap) minY() int {
	if cm.coordinates == nil || len(cm.coordinates) == 0 {
		panic("Cannot find min y for empty/nil map.")
	}
	minY := cm.coordinates[0].y
	for i := 1; i < len(cm.coordinates); i++ {
		y := cm.coordinates[i].y
		if y < minY {
			minY = y
		}
	}

	return minY
}

func (cm coordinateMap) maxX() int {
	if cm.coordinates == nil || len(cm.coordinates) == 0 {
		panic("Cannot find max x for empty/nil map.")
	}
	maxX := cm.coordinates[0].x
	for i := 1; i < len(cm.coordinates); i++ {
		x := cm.coordinates[i].x
		if x > maxX {
			maxX = x
		}
	}

	return maxX
}

func (cm coordinateMap) maxY() int {
	if cm.coordinates == nil || len(cm.coordinates) == 0 {
		panic("Cannot find max y for empty/nil map.")
	}
	maxY := cm.coordinates[0].y
	for i := 1; i < len(cm.coordinates); i++ {
		y := cm.coordinates[i].y
		if y > maxY {
			maxY = y
		}
	}

	return maxY
}

// If there is a point on the edge that is closest to a coordinate, that coordinate's area is unbounded
