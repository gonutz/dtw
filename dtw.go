package dtw

import "math"

func Match(a, b []float64) (matches [][2]int, cost float64) {
	costs := newMat(len(a)+1, len(b)+1)
	inf := math.Inf(1)
	for i := 0; i < costs.width; i++ {
		costs.set(i, 0, inf)
	}
	for i := 0; i < costs.height; i++ {
		costs.set(0, i, inf)
	}
	costs.set(0, 0, 0)

	for y := 1; y < costs.height; y++ {
		for x := 1; x < costs.width; x++ {
			d := a[x-1] - b[y-1]
			dist := d * d
			costs.set(x, y, dist+min(
				costs.get(x-1, y),
				costs.get(x-1, y-1),
				costs.get(x, y-1),
			))
		}
	}

	x, y := len(a)-1, len(b)-1
	for x >= 0 && y >= 0 {
		top := costs.get(x+1, y)
		topLeft := costs.get(x, y)
		left := costs.get(x, y+1)
		best := min(top, topLeft, left)

		matches = append(matches, [2]int{x, y})
		if best == topLeft || best == left {
			x--
		}
		if best == topLeft || best == top {
			y--
		}
	}

	// Reverse the path.
	n := len(matches)
	for i := 0; i < n/2; i++ {
		matches[i], matches[n-1-i] = matches[n-1-i], matches[i]
	}

	return matches, costs.get(len(a), len(b))
}

func newMat(w, h int) *mat {
	return &mat{
		width:  w,
		height: h,
		data:   make([]float64, w*h),
	}
}

type mat struct {
	width  int
	height int
	data   []float64
}

func (m *mat) set(x, y int, value float64) {
	m.data[x+y*m.width] = value
}

func (m *mat) get(x, y int) float64 {
	return m.data[x+y*m.width]
}

func min(a, b, c float64) float64 {
	if b < a {
		a, b = b, a
	}
	if c < a {
		a, c = c, a
	}
	return a
}
