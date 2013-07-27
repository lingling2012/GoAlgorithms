package p1_1

import (
	"testing"
)

func TestP1_1(t *testing.T) {
	const N = 10
	var cin = [][2]int{{3, 4},
		{4, 9},
		{8, 0},
		{2, 3},
		{5, 6},
		{2, 9},
		{5, 9},
		{7, 3},
		{4, 8},
		{5, 6},
		{0, 2},
		{6, 1},
	}

	FastFind(cin, N)
}

func TestE1_1_and_E1_4(t *testing.T) {
	const N = 10
	var cin = [][2]int{
		{0, 2},
		{1, 4},
		{2, 5},
		{3, 6},
		{0, 4},
		{6, 0},
		{1, 3},
	}

	FastFind(cin, N)
}