package day3

import "testing"

func TestParseClaim(t *testing.T) {
	cases := []struct {
		input  string
		number int
		left   int
		top    int
		right  int
		bottom int
		width  int
		height int
	}{
		{"#123 @ 3,2: 5x4", 123, 3, 2, 8, 6, 5, 4},
		{"#7 @ 30,2: 5x4", 7, 30, 2, 35, 6, 5, 4},
		{"#1002 @ 3,10: 50x12", 1002, 3, 10, 53, 22, 50, 12},
	}

	for _, c := range cases {
		claim := parseClaim(c.input)
		if claim.number != c.number {
			t.Errorf("ParseClaim for input %s had number %d but expected %d", c.input, claim.number, c.number)
		}
		if claim.left != c.left {
			t.Errorf("ParseClaim for input %s had left %d but expected %d", c.input, claim.left, c.left)
		}
		if claim.top != c.top {
			t.Errorf("ParseClaim for input %s had top %d but expected %d", c.input, claim.top, c.top)
		}
		if claim.right != c.right {
			t.Errorf("ParseClaim for input %s had right %d but expected %d", c.input, claim.right, c.right)
		}
		if claim.bottom != c.bottom {
			t.Errorf("ParseClaim for input %s had bottom %d but expected %d", c.input, claim.bottom, c.bottom)
		}
		if claim.width != c.width {
			t.Errorf("ParseClaim for input %s had width %d but expected %d", c.input, claim.width, c.width)
		}
		if claim.height != c.height {
			t.Errorf("ParseClaim for input %s had height %d but expected %d", c.input, claim.height, c.height)
		}
	}
}

func TestCalculateSquares(t *testing.T) {
	cases := []struct {
		claim    claim
		expected []square
	}{
		{claim{1, 1, 1, 3, 3, 2, 2}, []square{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
		{claim{1, 1, 1, 3, 2, 2, 1}, []square{{1, 1}, {2, 1}}},
		{claim{1, 1, 1, 2, 3, 1, 2}, []square{{1, 1}, {1, 2}}},
	}

	for _, c := range cases {
		squares := c.claim.calculateSquares()
		if len(squares) != len(c.expected) {
			t.Errorf("CalculateSquares had length of %d but expected %d", len(squares), len(c.expected))
		}

		for _, square := range c.expected {
			// Ordering doesn't matter
			found := false
			for _, calculatedSquare := range squares {
				if calculatedSquare.x == square.x && calculatedSquare.y == square.y {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("CalculateSquares expected square %v but was not found", square)
			}
		}
	}
}

func TestFindOverlappingSquares(t *testing.T) {
	claims := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	expected := 4
	overlapping := FindOverlappingSquares(claims)
	if overlapping != expected {
		t.Errorf("FindOverlappingSquares should find %d but got %d", expected, overlapping)
	}
}
