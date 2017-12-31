package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type direction int

const (
	up direction = iota
	right
	down
	left
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
	switch d {
	case up:
		return right
	case down:
		return left
	case right:
		return down
	case left:
		return up
	default:
		panic("whoops")
	}
}

func (d direction) TurnLeft() direction {
	switch d {
	case up:
		return left
	case down:
		return right
	case right:
		return up
	case left:
		return down
	default:
		panic("whoops")
	}
}

func (d direction) Reverse() direction {
	switch d {
	case up:
		return down
	case down:
		return up
	case right:
		return left
	case left:
		return right
	default:
		panic("whoops")
	}
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

type state int

const (
	clean state = iota
	weakened
	infected
	flagged
)

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
	states, position := parse("input.txt")

	direction := up
	iters := 10000000
	infections := 0

	for i := 0; i < iters; i++ {
		// print(infected, position)
		// fmt.Println()

		switch states[position] {
		case clean:
			direction = direction.TurnLeft()
			states[position] = weakened

		case weakened:
			states[position] = infected
			infections++

		case infected:
			direction = direction.TurnRight()
			states[position] = flagged

		case flagged:
			direction = direction.Reverse()
			states[position] = clean

		default:
			panic("whoops")
		}

		position.Move(direction)
	}

	// print(infected, position)
	fmt.Println(infections)
}

func parse(filename string) (map[point]state, point) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	result := map[point]state{}

	center := point{}
	center.x = len(lines[0]) / 2
	center.y = len(lines) / 2

	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				result[point{x, y}] = infected
			} else if char == '.' {
				result[point{x, y}] = clean
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
