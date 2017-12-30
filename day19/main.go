package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func parse(filename string) [][]byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")

	result := [][]byte{}
	for _, line := range lines {
		result = append(result, []byte(line))
	}

	return result
}

type position struct {
	x, y int
}

func (p position) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

type velocity position

func (p position) Move(v velocity) position {
	return position{x: p.x + v.x, y: p.y + v.y}
}

func (v velocity) IsZero() bool {
	return v.x == 0 && v.y == 0
}

func findStart(maze [][]byte) position {
	top := maze[0]
	for x := 0; x < len(top); x++ {
		if top[x] != ' ' {
			return position{x: x, y: 0}
		}
	}
	panic("nothing in the top row")
}

func canMoveHorizontal(maze [][]byte, x, y, dx int) bool {
	nx := x + dx
	if nx < 0 || nx >= len(maze[y]) {
		// That would be out of bounds.
		return false
	}

	c := maze[y][nx]
	switch c {
	case '-':
		fallthrough
	case '+':
		// It's definitely okay.
		return true

	case ' ':
		// It's definitely not okay.
		return false

	case '|':
		fallthrough
	default:
		// It might be okay. Can we collect it and keep moving?
		return canMoveHorizontal(maze, nx, y, dx)
	}
}

func canMoveVertical(maze [][]byte, x, y, dy int) bool {
	ny := y + dy
	if ny < 0 || ny >= len(maze) {
		// That would be out of bounds.
		return false
	}

	c := maze[ny][x]
	switch c {
	case '|':
		fallthrough
	case '+':
		// It's definitely okay.
		return true

	case ' ':
		// It's definitely not okay.
		return false

	case '-':
		fallthrough
	default:
		// It might be okay. Can we collect it and keep moving?
		return canMoveVertical(maze, x, ny, dy)
	}
}

func newDirection(maze [][]byte, cur position, vel velocity) velocity {
	if vel.x == 0 {
		// Currently moving up/down.
		if canMoveHorizontal(maze, cur.x, cur.y, -1) {
			// We can move left, let's do that.
			return velocity{x: -1, y: 0}
		}
		if canMoveHorizontal(maze, cur.x, cur.y, 1) {
			// We can move right, let's do that.
			return velocity{x: 1, y: 0}
		}
	} else {
		// Currently moving left/right.
		if canMoveVertical(maze, cur.x, cur.y, -1) {
			// We can move up, let's do that.
			return velocity{x: 0, y: -1}
		}
		if canMoveVertical(maze, cur.x, cur.y, 1) {
			// We can move down, let's do that.
			return velocity{x: 0, y: 1}
		}
	}

	panic(fmt.Sprintf("dead end at %v!", cur))
}

func iLikeToMoveItMoveIt(maze [][]byte) (position, []byte, int) {
	cur := findStart(maze)
	vel := velocity{x: 0, y: 1}

	chars := []byte{}
	steps := 0

	for {
		c := maze[cur.y][cur.x]

		switch c {
		case ' ':
			return cur, chars, steps

		case '|':
			fallthrough
		case '-':
			// Keep going in the direction we were going.

		case '+':
			// Make a turn.
			vel = newDirection(maze, cur, vel)

		default:
			// Collect the letter and keep moving.
			chars = append(chars, c)
		}

		cur = cur.Move(vel)
		steps++
	}
}

func main() {
	// maze := parse("test.txt")
	maze := parse("input.txt")

	cur, chars, steps := iLikeToMoveItMoveIt(maze)

	fmt.Println(cur)
	fmt.Println(string(chars))
	fmt.Println(steps)
}
