package day3

import (
	"regexp"
	"strconv"
)

type claim struct {
	number int
	left   int
	top    int
	right  int
	bottom int
	width  int
	height int
}

type square struct {
	x int
	y int
}

var claimRegex = regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

// ParseClaim reads a claim in the format "#NUM @ L,T: WxH",
// e.g. #123 @ 3,2: 5x4, which means claim number 123 is 3 units
// from the left, 2 from the top, and has a size of 5 units width
// by 4 units height.
func parseClaim(input string) claim {
	parsed := claimRegex.FindStringSubmatch(input)[1:]
	var parts []int
	for _, part := range parsed {
		intPart, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		parts = append(parts, intPart)
	}

	right := parts[1] + parts[3]
	bottom := parts[2] + parts[4]
	return claim{parts[0], parts[1], parts[2], right, bottom, parts[3], parts[4]}
}

func (c claim) calculateSquares() []square {
	var squares []square
	for i := c.left; i < c.right; i++ {
		for j := c.top; j < c.bottom; j++ {
			squares = append(squares, square{i, j})
		}
	}

	return squares
}

// FindOverlappingSquares takes a slice of claim strings, parses them and finds
// how many of squares they overlap on.
func FindOverlappingSquares(input []string) int {
	foundSquares := make(map[square]struct{})
	overlappingSquares := make(map[square]struct{})

	for _, s := range input {
		claim := parseClaim(s)
		squares := claim.calculateSquares()
		for _, square := range squares {
			_, ok := foundSquares[square]
			if ok {
				overlappingSquares[square] = struct{}{}
			} else {
				foundSquares[square] = struct{}{}
			}
		}
	}

	return len(overlappingSquares)
}
