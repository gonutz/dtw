package dtw_test

import (
	"testing"

	"github.com/gonutz/check"
	"github.com/gonutz/dtw"
)

func TestExactSubPatter(t *testing.T) {
	//             0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
	a := []float64{1, 1, 1, 1, 2, 3, 4, 5, 4, 3, 2, 1, 1, 1, 1, 1}
	//             0  1  2  3  4  5  6
	b := []float64{2, 3, 4, 5, 4, 3, 2}
	matches, cost := dtw.Match(a, b)
	check.Eq(t, matches, [][2]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 1},
		{6, 2},
		{7, 3},
		{8, 4},
		{9, 5},
		{10, 6},
		{11, 6},
		{12, 6},
		{13, 6},
		{14, 6},
		{15, 6},
	})
	check.Eq(t, cost, 9)
}

func TestStretchedPattern(t *testing.T) {
	//             0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
	a := []float64{1, 1, 1, 1, 2, 3, 4, 5, 4, 3, 2, 1, 1, 1, 1, 1}
	//             0  1  2  3  4  5  6  7  8  9 10 11 12 13 14
	b := []float64{2, 2, 3, 3, 4, 4, 4, 5, 5, 4, 4, 3, 3, 2, 2}
	matches, cost := dtw.Match(a, b)
	check.Eq(t, matches, [][2]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 1},
		{5, 2},
		{5, 3},
		{6, 4},
		{6, 5},
		{6, 6},
		{7, 7},
		{7, 8},
		{8, 9},
		{8, 10},
		{9, 11},
		{9, 12},
		{10, 13},
		{11, 13},
		{12, 13},
		{13, 13},
		{14, 13},
		{15, 14},
	})
	check.Eq(t, cost, 9)
}

func TestAmplifiedPattern(t *testing.T) {
	//             0  1  2  3  4  5  6
	a := []float64{1, 1, 2, 3, 2, 1, 1}
	//             0  1  2  3  4  5  6  7
	b := []float64{0, 1, 1, 3, 5, 3, 1, 1}
	matches, cost := dtw.Match(a, b)
	check.Eq(t, matches, [][2]int{
		{0, 0},
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
	})
	check.Eq(t, cost, 7)
}
