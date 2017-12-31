package main

import "fmt"

type visitor interface {
	visit(n, x, y int) bool
}

func visit(v visitor) (x, y int) {
	n := 1
	x = 0
	y = 0

	if v.visit(n, x, y) {
		return
	}

	distance := 1

	for {
		// Move `distance` units right.
		for i := 0; i < distance; i++ {
			x++
			n++
			if v.visit(n, x, y) {
				return
			}
		}

		// Move `distance` units up.
		for i := 0; i < distance; i++ {
			y++
			n++
			if v.visit(n, x, y) {
				return
			}
		}

		distance++

		// Move `distance` units left.
		for i := 0; i < distance; i++ {
			x--
			n++
			if v.visit(n, x, y) {
				return
			}
		}

		// Move `distance` units down.
		for i := 0; i < distance; i++ {
			y--
			n++
			if v.visit(n, x, y) {
				return
			}
		}

		distance++
	}
}

type finder struct {
	target int
}

func (f finder) visit(n, x, y int) bool {
	return n == f.target
}

type point struct {
	x, y int
}

type memtester struct {
	target int
	state  map[point]int
}

func (m *memtester) visit(n, x, y int) bool {
	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			sum += m.state[point{x + i, y + j}]
		}
	}
	m.state[point{x, y}] = sum

	return sum > m.target
}

func main() {
	// Part one.
	{
		find := finder{target: 368078}
		x, y := visit(find)

		if x < 0 {
			x = -x
		}
		if y < 0 {
			y = -y
		}
		fmt.Printf("moves: %v\n", x+y)
	}

	// Part two.
	{
		test := memtester{target: 368078, state: make(map[point]int)}
		test.state[point{0, 0}] = 1
		x, y := visit(&test)

		fmt.Printf("result: %v\n", test.state[point{x, y}])
	}
}
