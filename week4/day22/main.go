package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type direction int

const (
	up    = direction(0)
	right = direction(1)
	down  = direction(2)
	left  = direction(3)
)

func (d direction) DX() int {
	switch d {
	case left:
		return -1
	case right:
		return 1
	default:
		return 0
	}
}

func (d direction) DY() int {
	switch d {
	case up:
		return -1
	case down:
		return 1
	default:
		return 0
	}
}

func (d direction) TurnRight() direction {
	return (d + 1) % 4
}

func (d direction) TurnLeft() direction {
	if d == 0 {
		return direction(3)
	}
	return d - 1
}

func (d direction) String() string {
	switch d {
	case up:
		return "up"
	case down:
		return "down"
	case left:
		return "left"
	case right:
		return "right"
	default:
		panic("whoops")
	}
}

type point struct {
	x, y int
}

func (p *point) Move(dir direction) {
	p.x += dir.DX()
	p.y += dir.DY()
}

func (p *point) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func main() {
	// infected, position := parse("test.txt")
	infected, position := parse("input.txt")

	direction := up
	iters := 10000
	infections := 0

	for i := 0; i < iters; i++ {
		// print(infected, position)
		// fmt.Println()

		if infected[position] {
			direction = direction.TurnRight()
			infected[position] = false
		} else {
			direction = direction.TurnLeft()
			infected[position] = true
			infections++
		}

		position.Move(direction)
	}

	// print(infected, position)
	fmt.Println(infections)
}

func parse(filename string) (map[point]bool, point) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	result := map[point]bool{}

	center := point{}
	center.x = len(lines[0]) / 2
	center.y = len(lines) / 2

	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				result[point{x, y}] = true
			} else if char == '.' {
				result[point{x, y}] = false
			} else {
				panic("whoops")
			}
		}
	}

	return result, center
}

func print(infected map[point]bool, position point) {
	minx := position.x
	miny := position.y
	maxx := position.x
	maxy := position.y

	for pt := range infected {
		if pt.x < minx {
			minx = pt.x
		}
		if pt.x > maxx {
			maxx = pt.x
		}
		if pt.y < miny {
			miny = pt.y
		}
		if pt.y > maxy {
			maxy = pt.y
		}
	}

	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if infected[point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

			if x == position.x && y == position.y {
				fmt.Print("]")
			} else if x+1 == position.x && y == position.y {
				fmt.Print("[")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
